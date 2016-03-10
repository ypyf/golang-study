// defines the namespace
namespace go wifi.session.rpc
namespace php wifi.session.rpc
namespace java wifi.session.rpc

const string VERSION = "1.0.0"

struct  SessionContext {
    1: required string aes_key,
    2: required string uuid,
    3: required string token,
    4: required string package_name,
    5: required string uid,
    6: required string imei,
    7: required string info,
    8: required string version,
}

exception InvalidSession {
}

exception InvalidArguments {
    1: required string why
}

service SessionManager {
    string create_session(1: required SessionContext ctx),
    SessionContext get_session(1: required string sid) throws (1:InvalidSession ouch),
}
