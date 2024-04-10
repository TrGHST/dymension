package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	porttypes "github.com/cosmos/ibc-go/v6/modules/core/05-port/types"
	commontypes "github.com/dymensionxyz/dymension/v3/x/common/types"

	osmoutils "github.com/osmosis-labs/osmosis/v15/osmoutils"
	"github.com/tendermint/tendermint/libs/log"
)

// FinalizeRollappPackets finalizes the packets for the given rollapp until the given height which is
// the end height of the latest finalized state
func (k Keeper) FinalizeRollappPackets(ctx sdk.Context, ibc porttypes.IBCModule, rollappID string, stateEndHeight uint64) error {
	rollappPendingPackets := k.ListRollappPackets(ctx, PendingByRollappIDByMaxHeight(rollappID, stateEndHeight))
	if len(rollappPendingPackets) == 0 {
		return nil
	}
	logger := ctx.Logger().With("module", "DelayedAckMiddleware")
	// Get the packets for the rollapp until height
	logger.Debug("finalizing IBC rollapp packets",
		"rollappID", rollappID,
		"state end height", stateEndHeight,
		"num packets", len(rollappPendingPackets))
	for _, rollappPacket := range rollappPendingPackets {
		if err := k.finalizeRollappPacket(ctx, ibc, rollappID, logger, rollappPacket); err != nil {
			return fmt.Errorf("finalize rollapp packet: %w", err)
		}
	}
	return nil
}

type wrappedFunc func(ctx sdk.Context) error

func (k Keeper) finalizeRollappPacket(
	ctx sdk.Context,
	ibc porttypes.IBCModule,
	rollappID string,
	logger log.Logger,
	rollappPacket commontypes.RollappPacket,
) (err error) {
	logContext := []interface{}{
		"rollappID", rollappID,
		"sequence", rollappPacket.Packet.Sequence,
		"source channel", rollappPacket.Packet.SourceChannel,
		"destination channel", rollappPacket.Packet.DestinationChannel,
		"type", rollappPacket.Type,
	}

	switch rollappPacket.Type {
	case commontypes.RollappPacket_ON_RECV:
		err = osmoutils.ApplyFuncIfNoError(ctx, k.onRecvPacket(rollappPacket, ibc))
	case commontypes.RollappPacket_ON_ACK:
		err = osmoutils.ApplyFuncIfNoError(ctx, k.onAckPacket(rollappPacket, ibc))
	case commontypes.RollappPacket_ON_TIMEOUT:
		err = osmoutils.ApplyFuncIfNoError(ctx, k.onTimeoutPacket(rollappPacket, ibc))
	default:
		logger.Error("Unknown rollapp packet type", logContext...)
	}
	// Update the packet with the error
	if err != nil {
		rollappPacket.Error = err.Error()
	}
	// Update status to finalized
	_, err = k.UpdateRollappPacketWithStatus(ctx, rollappPacket, commontypes.Status_FINALIZED)
	if err != nil {
		// If we failed finalizing the packet we return an error to abort the end blocker otherwise it's
		// invariant breaking
		return err
	}

	logger.Debug("finalized IBC rollapp packet", logContext...)
	return
}

func (k Keeper) onRecvPacket(rollappPacket commontypes.RollappPacket, ibc porttypes.IBCModule) wrappedFunc {
	return func(ctx sdk.Context) (err error) {
		ack := ibc.OnRecvPacket(ctx, *rollappPacket.Packet, rollappPacket.Relayer)
		// If async, return
		if ack == nil {
			return
		}

		// Write the acknowledgement to the chain only if it is synchronous
		var chanCap *capabilitytypes.Capability
		_, chanCap, err = k.LookupModuleByChannel(
			ctx,
			rollappPacket.Packet.DestinationPort,
			rollappPacket.Packet.DestinationChannel,
		)
		if err != nil {
			return
		}
		err = k.WriteAcknowledgement(ctx, chanCap, rollappPacket.Packet, ack)
		return
	}
}

func (k Keeper) onAckPacket(rollappPacket commontypes.RollappPacket, ibc porttypes.IBCModule) wrappedFunc {
	return func(ctx sdk.Context) (err error) {
		return ibc.OnAcknowledgementPacket(
			ctx,
			*rollappPacket.Packet,
			rollappPacket.Acknowledgement,
			rollappPacket.Relayer,
		)
	}
}

func (k Keeper) onTimeoutPacket(rollappPacket commontypes.RollappPacket, ibc porttypes.IBCModule) wrappedFunc {
	return func(ctx sdk.Context) (err error) {
		return ibc.OnTimeoutPacket(ctx, *rollappPacket.Packet, rollappPacket.Relayer)
	}
}
