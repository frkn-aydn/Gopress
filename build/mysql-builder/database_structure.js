const database =  process.env.MYSQL_DATABASE || "gopress"; // [todo] düzelt burayı
if(!database) return throwError("Mysql database not found!")

function throwError(message){
    console.log("\x1b[31m", message)
    return process.exit()
}

module.exports = {
    database : database,
    tables : [
        {
            name : "users",
            rows : {
                id : "INT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY",
                email : "VARCHAR(200) NULL, UNIQUE INDEX `email_UNIQUE` (`email` ASC)",
                message : "VARCHAR(500)",
                name : "VARCHAR(100)"
            }
        }
    ]
}