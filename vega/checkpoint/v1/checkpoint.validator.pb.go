// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vega/checkpoint/v1/checkpoint.proto

package v1

import (
	fmt "fmt"
	math "math"

	_ "code.vegaprotocol.io/sdk-golang/vega"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CheckpointState) Validate() error {
	return nil
}
func (this *Checkpoint) Validate() error {
	return nil
}
func (this *AssetEntry) Validate() error {
	if this.AssetDetails != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AssetDetails); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AssetDetails", err)
		}
	}
	return nil
}
func (this *Assets) Validate() error {
	for _, item := range this.Assets {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Assets", err)
			}
		}
	}
	return nil
}
func (this *AssetBalance) Validate() error {
	return nil
}
func (this *Collateral) Validate() error {
	for _, item := range this.Balances {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Balances", err)
			}
		}
	}
	return nil
}
func (this *NetParams) Validate() error {
	for _, item := range this.Params {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
			}
		}
	}
	return nil
}
func (this *Proposals) Validate() error {
	for _, item := range this.Proposals {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Proposals", err)
			}
		}
	}
	return nil
}
func (this *DelegateEntry) Validate() error {
	return nil
}
func (this *Delegate) Validate() error {
	for _, item := range this.Active {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Active", err)
			}
		}
	}
	for _, item := range this.Pending {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Pending", err)
			}
		}
	}
	return nil
}
func (this *Block) Validate() error {
	return nil
}
func (this *Rewards) Validate() error {
	for _, item := range this.Rewards {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Rewards", err)
			}
		}
	}
	return nil
}
func (this *RewardPayout) Validate() error {
	for _, item := range this.RewardsPayout {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RewardsPayout", err)
			}
		}
	}
	return nil
}
func (this *PendingRewardPayout) Validate() error {
	for _, item := range this.PartyAmount {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PartyAmount", err)
			}
		}
	}
	return nil
}
func (this *PartyAmount) Validate() error {
	return nil
}
func (this *PendingKeyRotation) Validate() error {
	return nil
}
func (this *KeyRotations) Validate() error {
	for _, item := range this.PendingKeyRotations {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PendingKeyRotations", err)
			}
		}
	}
	return nil
}
