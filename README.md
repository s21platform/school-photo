# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [school.proto](#school-proto)
    - [Campus](#-Campus)
    - [CampusUuidIn](#-CampusUuidIn)
    - [CampusesOut](#-CampusesOut)
    - [Empty](#-Empty)
    - [PeerRequest](#-PeerRequest)
    - [PeerResponse](#-PeerResponse)
    - [SchoolLoginRequest](#-SchoolLoginRequest)
    - [SchoolLoginResponse](#-SchoolLoginResponse)
    - [Tribe](#-Tribe)
    - [TribesOut](#-TribesOut)
  
    - [SchoolService](#-SchoolService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="school-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## school.proto



<a name="-Campus"></a>

### Campus
Сообщение для выходных данных


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| campusUuid | [string](#string) |  | Uuid кампуса |
| shortName | [string](#string) |  | Сокращенное название кампуса |
| fullName | [string](#string) |  | Полное название кампуса |






<a name="-CampusUuidIn"></a>

### CampusUuidIn
Запрос на получение всех трайбов кампуса


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| campus_uuid | [string](#string) |  | Uuid кампуса, который призходит от community сервиса |






<a name="-CampusesOut"></a>

### CampusesOut
Ответ на запрос на получение всех кампусов


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| campuses | [Campus](#Campus) | repeated |  |






<a name="-Empty"></a>

### Empty
Запрос на получение всех кампусов






<a name="-PeerRequest"></a>

### PeerRequest
Запрос на получение списка пиров


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| campusUuid | [string](#string) |  | Uuid кампуса |
| limit | [int64](#int64) |  | Kоличество записей для возвращения |
| offset | [int64](#int64) |  | Смещение |






<a name="-PeerResponse"></a>

### PeerResponse
Ответ на запрос на получение списка пиров


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [string](#string) | repeated |  |






<a name="-SchoolLoginRequest"></a>

### SchoolLoginRequest
Запрос на авторизацию


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  | email от edu.21-school.ru формата nickname@student.21-school.ru |
| password | [string](#string) |  | Пароль пользователя от edu.21-school.ru |






<a name="-SchoolLoginResponse"></a>

### SchoolLoginResponse
Ответ на запрос авторизации


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  | Авторизационный токен JWT |






<a name="-Tribe"></a>

### Tribe
Сообщение для выходных данных


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | id тайба |
| name | [string](#string) |  | Название трайба |






<a name="-TribesOut"></a>

### TribesOut
Ответ на запрос на получение всех трайбов кампуса


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tribes | [Tribe](#Tribe) | repeated | Список трайбов |





 

 

 


<a name="-SchoolService"></a>

### SchoolService
Этот сервис предоставляет методы для запросов на школьную платформу.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Login | [.SchoolLoginRequest](#SchoolLoginRequest) | [.SchoolLoginResponse](#SchoolLoginResponse) | Метод для получения токена sberclass, которым будут подписаны запросы к платформе |
| GetCampuses | [.Empty](#Empty) | [.CampusesOut](#CampusesOut) | Методы для получения данных с API школы |
| GetTribesByCampusUuid | [.CampusUuidIn](#CampusUuidIn) | [.TribesOut](#TribesOut) |  |
| GetPeers | [.PeerRequest](#PeerRequest) | [.PeerResponse](#PeerResponse) | Метод для получения списка пиров |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

