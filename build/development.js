// Dependencies
const ora = require('ora');
const webpack = require("webpack");
const chalk = require("chalk")
const commandLineArgs = require('command-line-args')

const optionDefinitions = [{
    name: 'src',
    type: String,
    multiple: false
}]

const options = commandLineArgs(optionDefinitions);

if(!options.src){
    console.log("Please enter at least one file...")
    process.exit(1)
}

// Config files
const wpConfig = require("./webpack.dev.conf")


const spinner = ora('Building for production...').start();

webpack(wpConfig(options.src.trim()), function (err, stats) {
    spinner.stop()
    if (err) throw err
    process.stdout.write(stats.toString({
        colors: true,
        modules: false,
        children: false,
        chunks: false,
        chunkModules: false
    }) + '\n\n')
    console.log(chalk.cyan('  Build complete.\n'))
    console.log(chalk.yellow(
        '  Tip: built files are meant to be served over an HTTP server.\n' +
        '  Opening index.html over file:// won\'t work.\n'
    ))
})