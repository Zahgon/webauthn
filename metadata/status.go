package metadata

// ValidateStatusReports checks a list of [StatusReport] structs against a list of desired and undesired [AuthenticatorStatus]
// values. If the reports contain all of the desired and none of the undesired status reports then no error is returned
// otherwise an error describing the issue is returned.
//
//nolint:gocyclo
func ValidateStatusReports(reports []StatusReport, desired, undesired []AuthenticatorStatus) (err error) {
	_ = "STUB: not implemented"
	return nil
}
