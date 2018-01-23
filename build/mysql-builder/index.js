const dotenv = require(`dotenv`).config()
const mysql = require('mysql');
const structure = require("./database_structure");
const dbConfig = {
    host     : 'localhost',
    user     : 'root', // for example...
    password : 'password' // for example...
}
let connection = mysql.createConnection(dbConfig);

connection.connect();
connection.query(`CREATE DATABASE IF NOT EXISTS ${structure.database} DEFAULT CHARACTER SET utf8;`, function (err, results, fields) {
    if (err){
        console.log(err)
        throw err;
    }
    dbConfig.database = structure.database;
    connection = mysql.createConnection(dbConfig)
    structure.tables.forEach(table => {
        const rows = table.rows;
        const rowQuery = []
        for (let name in rows) {
            if (rows.hasOwnProperty(name)) {
                rowQuery.push(name + " " + rows[name])
            }
        }
        connection.query(`CREATE TABLE IF NOT EXISTS ${table.name} (${rowQuery.join(", ")});`, function (error, results, fields) {
            if (error) throw error;
        });
    })
});
connection.end();
console.log("Database and tables generated successfuly.");