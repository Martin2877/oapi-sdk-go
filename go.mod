module github.com/martin2877/oapi-sdk-go/v3

go 1.13

//
//replace (
//	github.com/martin2877/oapi-sdk-go/v3/core => ./core
//	github.com/martin2877/oapi-sdk-go/v3/event => ./event
//	github.com/martin2877/oapi-sdk-go/v3/service/acs => ./service/acs
//	github.com/martin2877/oapi-sdk-go/v3/service/admin => ./service/admin
//	github.com/martin2877/oapi-sdk-go/v3/service/application => ./service/application
//	github.com/martin2877/oapi-sdk-go/v3/service/approval => ./service/approval
//	github.com/martin2877/oapi-sdk-go/v3/service/attendance => ./service/attendance
//	github.com/martin2877/oapi-sdk-go/v3/service/baike => ./service/baike
//	github.com/martin2877/oapi-sdk-go/v3/service/bitable => ./service/bitable
//	github.com/martin2877/oapi-sdk-go/v3/service/calendar => ./service/calendar
//	github.com/martin2877/oapi-sdk-go/v3/service/contact => ./service/contact
//	github.com/martin2877/oapi-sdk-go/v3/service/docx => ./service/docx
//	github.com/martin2877/oapi-sdk-go/v3/service/drive => ./service/drive
//	github.com/martin2877/oapi-sdk-go/v3/service/ehr => ./service/ehr
//	github.com/martin2877/oapi-sdk-go/v3/service/event => ./service/event
//	github.com/martin2877/oapi-sdk-go/v3/service/human_authentication => ./service/human_authentication
//	github.com/martin2877/oapi-sdk-go/v3/service/im => ./service/im
//	github.com/martin2877/oapi-sdk-go/v3/service/mail => ./service/mail
//	github.com/martin2877/oapi-sdk-go/v3/service/meeting_room => ./service/meeting_room
//	github.com/martin2877/oapi-sdk-go/v3/service/optical_char_recognition => ./service/optical_char_recognition
//	github.com/martin2877/oapi-sdk-go/v3/service/passport => ./service/passport
//	github.com/martin2877/oapi-sdk-go/v3/service/search => ./service/search
//	github.com/martin2877/oapi-sdk-go/v3/service/sheets => ./service/sheets
//	github.com/martin2877/oapi-sdk-go/v3/service/speech_to_text => ./service/speech_to_text
//	github.com/martin2877/oapi-sdk-go/v3/service/task => ./service/task
//	github.com/martin2877/oapi-sdk-go/v3/service/tenant => ./service/tenant
//	github.com/martin2877/oapi-sdk-go/v3/service/translation => ./service/translation
//	github.com/martin2877/oapi-sdk-go/v3/service/vc => ./service/vc
//	github.com/martin2877/oapi-sdk-go/v3/service/wiki => ./service/wiki
//
//)

require (
	github.com/gogo/protobuf v1.3.2
	github.com/gorilla/websocket v1.5.0
)
