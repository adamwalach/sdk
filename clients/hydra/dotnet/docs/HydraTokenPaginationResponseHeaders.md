# Ory.Hydra.Client.Model.HydraTokenPaginationResponseHeaders
The `Link` HTTP header contains multiple links (`first`, `next`, `last`, `previous`) formatted as: `<https://{project-slug}.projects.oryapis.com/admin/clients?limit={limit}&offset={offset}>; rel=\"{page}\"`  For details on pagination please head over to the [pagination documentation](https://www.ory.sh/docs/ecosystem/api-design#pagination).

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | **string** | The Link HTTP Header  The &#x60;Link&#x60; header contains a comma-delimited list of links to the following pages:  first: The first page of results. next: The next page of results. prev: The previous page of results. last: The last page of results.  Pages are omitted if they do not exist. For example, if there is no next page, the &#x60;next&#x60; link is omitted. Examples:  &lt;/clients?limit&#x3D;5&amp;offset&#x3D;0&gt;; rel&#x3D;\&quot;first\&quot;,&lt;/clients?limit&#x3D;5&amp;offset&#x3D;15&gt;; rel&#x3D;\&quot;next\&quot;,&lt;/clients?limit&#x3D;5&amp;offset&#x3D;5&gt;; rel&#x3D;\&quot;prev\&quot;,&lt;/clients?limit&#x3D;5&amp;offset&#x3D;20&gt;; rel&#x3D;\&quot;last\&quot; | [optional] 
**XTotalCount** | **long** | The X-Total-Count HTTP Header  The &#x60;X-Total-Count&#x60; header contains the total number of items in the collection. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

