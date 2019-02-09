package reisapi-ruter

const dateFormat = "20060102"

type HeartbeatRequest struct {

}

type HeartbeatResponse struct {
	OtpVersion string,
	Otp	bool,
	TransitServiceStarts string,	
	TransitServiceEnds string,
	ReisVersion string,
	OtpResponseTime int32,
	Sql bool,
	SqlResponseTime string,
	Els bool,
	ElsVersion string,
	ElsResponsetime int32
}
