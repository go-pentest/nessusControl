package nessusProcessor

// PolicyViolationMatchCriteria holds what criteria should be checked when checking for a
// policy violation.
type PolicyViolationMatchCriteria struct {
	ExternallyAccessible             bool
	IgnoreViolationWithCriteriaMatch bool
	PreviousViolationCheck           bool
	CountIf                          string
	DescriptionRegexp                []string
	NotDescriptionRegexp             []string
	PluginID                         int
	Ports                            []int
	OrganizationIDs                  []int
	RegionIDs                        []int
}

// FalsePositiveMatchCriteria holds what criteria should be checked when
// checking for a false positive.
type FalsePositiveMatchCriteria struct {
	PluginID                 int
	Ports                    []int
	Protocol                 string
	DescriptionRegexp        []string
	IsIPWithinExternalRegion bool
	SQLSolarisCheck          bool
}
