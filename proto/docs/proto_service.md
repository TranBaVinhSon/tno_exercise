# Protocol Documentation
<a name="top"/>

## Table of Contents

- [service.proto](#service.proto)
    - [GetBalanceRequest](#services.GetBalanceRequest)
    - [GetBalanceResponse](#services.GetBalanceResponse)
    - [GetTransactionsRequest](#services.GetTransactionsRequest)
    - [GetTransactionsResponse](#services.GetTransactionsResponse)
    - [SendCoinRequest](#services.SendCoinRequest)
    - [SendCoinResponse](#services.SendCoinResponse)
    - [Transaction](#services.Transaction)
  
  
  
    - [Wallet](#services.Wallet)
  

- [Scalar Value Types](#scalar-value-types)



<a name="service.proto"/>
<p align="right"><a href="#top">Top</a></p>

## service.proto



<a name="services.GetBalanceRequest"/>

### GetBalanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  | ユーザーID |






<a name="services.GetBalanceResponse"/>

### GetBalanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| balance | [string](#string) |  | 残高 |






<a name="services.GetTransactionsRequest"/>

### GetTransactionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [uint64](#uint64) |  |  |






<a name="services.GetTransactionsResponse"/>

### GetTransactionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transactions | [Transaction](#services.Transaction) | repeated |  |






<a name="services.SendCoinRequest"/>

### SendCoinRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| from_user_id | [uint64](#uint64) |  |  |
| to_user_id | [uint64](#uint64) |  |  |
| amount | [string](#string) |  |  |






<a name="services.SendCoinResponse"/>

### SendCoinResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_id | [string](#string) |  |  |






<a name="services.Transaction"/>

### Transaction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| category | [string](#string) |  |  |
| abandoned | [string](#string) |  |  |
| received_account | [string](#string) |  |  |
| received_address | [string](#string) |  |  |
| amount | [string](#string) |  |  |





 

 

 


<a name="services.Wallet"/>

### Wallet


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetBalance | [GetBalanceRequest](#services.GetBalanceRequest) | [GetBalanceResponse](#services.GetBalanceRequest) | ログイン情報取得 |
| SendCoin | [SendCoinRequest](#services.SendCoinRequest) | [SendCoinResponse](#services.SendCoinRequest) |  |
| GetTransactions | [GetTransactionsRequest](#services.GetTransactionsRequest) | [GetTransactionsResponse](#services.GetTransactionsRequest) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

