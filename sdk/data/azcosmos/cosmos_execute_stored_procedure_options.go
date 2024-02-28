package azcosmos

// ExecuteStoredProcedureOptions includes options for executing a stored procedure.
type ExecuteStoredProcedureOptions struct {
	EnableScriptLogging bool
}

func (options *ExecuteStoredProcedureOptions) toHeaders() *map[string]string {
	headers := make(map[string]string)

	if options.EnableScriptLogging {
		headers[cosmosHeaderEnableScriptLogging] = "true"
	}

	return &headers
}
