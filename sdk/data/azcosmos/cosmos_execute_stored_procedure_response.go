package azcosmos

import (
	"net/http"

	azruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

type ExecuteStoredProcedureResponse struct {
	Value     []byte
	ScriptLog string
	Response
}

func newExecuteStoredProcedureResponse(resp *http.Response) (ExecuteStoredProcedureResponse, error) {
	response := ExecuteStoredProcedureResponse{
		Response: newResponse(resp),
	}

	defer resp.Body.Close()
	body, err := azruntime.Payload(resp)
	if err != nil {
		return response, err
	}
	response.Value = body

	logResults := resp.Header.Get(cosmosHeaderScriptLogResults)
	if logResults != "" {
		response.ScriptLog = logResults // TODO unescape?
	}

	return response, nil
}
