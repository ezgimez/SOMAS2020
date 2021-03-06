package baseclient

import (
	"github.com/SOMAS2020/SOMAS2020/internal/common/roles"
	"github.com/SOMAS2020/SOMAS2020/internal/common/rules"
)

// CommonPoolResourceRequest is called by the President in IIGO to
// request an allocation of resources from the common pool.
func (c *BaseClient) CommonPoolResourceRequest() int {
	// TODO: Implement needs based resource request.
	return 20
}

// ResourceReport is an island's self-report of its own resources.
func (c *BaseClient) ResourceReport() int {
	return c.clientGameState.ClientInfo.Resources
}

// RuleProposal is called by the President in IIGO to propose a
// rule to be voted on.
func (c *BaseClient) RuleProposal() string {
	allRules := rules.AvailableRules
	for k, _ := range allRules {
		return k
	}
	return ""
}

// GetClientPresidentPointer is called by IIGO to get the client's implementation of the President Role
func (c *BaseClient) GetClientPresidentPointer() roles.President {
	return nil
}

// GetClientJudgePointer is called by IIGO to get the client's implementation of the Judge Role
func (c *BaseClient) GetClientJudgePointer() roles.Judge {
	return nil
}

// GetClientSpeakerPointer is called by IIGO to get the client's implementation of the Speaker Role
func (c *BaseClient) GetClientSpeakerPointer() roles.Speaker {
	return nil
}

func (c *BaseClient) GetTaxContribution() int {
	// TODO: Implement common pool contribution greater than or equal to tax.
	return 0
}

func (c *BaseClient) RequestAllocation() int {
	// TODO: Implement request equal to the allocation permitted by President.
	return 0
}
