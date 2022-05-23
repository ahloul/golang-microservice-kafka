conn = new Mongo();
db = conn.getDB("reports")
db.createUser(
    {
        user: "reports_user",
        pwd: "root",
        roles: [
            {
                role: "dbOwner",
                db: "reports"
            }
        ]
    }
);
db.createCollection('users', { capped: false });
db.createCollection('purchases', { capped: false });
db.createCollection('redeems', { capped: false });
