# ory_hydra_client.model.TokenPaginationResponseHeaders

## Load the model package
```dart
import 'package:ory_hydra_client/api.dart';
```

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**link** | **String** | The Link HTTP Header  The `Link` header contains a comma-delimited list of links to the following pages:  first: The first page of results. next: The next page of results. prev: The previous page of results. last: The last page of results.  Pages are omitted if they do not exist. For example, if there is no next page, the `next` link is omitted. Examples:  </clients?limit=5&offset=0>; rel=\"first\",</clients?limit=5&offset=15>; rel=\"next\",</clients?limit=5&offset=5>; rel=\"prev\",</clients?limit=5&offset=20>; rel=\"last\" | [optional] 
**xTotalCount** | **int** | The X-Total-Count HTTP Header  The `X-Total-Count` header contains the total number of items in the collection. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


