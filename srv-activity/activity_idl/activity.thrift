namespace go activity

include "product.thrift"

struct Activity {
    1: i64 id,
    2: optional string name,
    3: optional string desc,
    4: optional string start_time,
    5: optional string end_time,
}

service ActivityService {
    list <product.Product> GetActivities(Activity a),
    Activity GetActivityInfo (Activity a)
}