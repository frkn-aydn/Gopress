// Dependencies
const webpack = require("webpack");
const HtmlWebpackPlugin = require('html-webpack-plugin');
const ExtractTextPlugin = require("extract-text-webpack-plugin");
const CopyWebpackPlugin = require('copy-webpack-plugin')
const OptimizeCSSPlugin = require('optimize-css-assets-webpack-plugin')
const SWPrecacheWebpackPlugin = require('sw-precache-webpack-plugin')
const UglifyJsPlugin = require('uglifyjs-webpack-plugin');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const OfflinePlugin = require('offline-plugin');
const HtmlWebpackInlineSourcePlugin = require('html-webpack-inline-source-plugin');

const path = require("path");
const fs = require("fs")
//  Useful functions
const utils = require("./utils");

// Getting node evn.
const env = '"development"'
const extractLess = new ExtractTextPlugin({
	filename: "css/[name].css"
});


function webpackConfigGenerator(file) {
    // Generating webpack settings...
    const webpackConfig = {
        entry: {
            demo: path.resolve(__dirname, "../client/js/" + file + ".js")
        },
        output: {
            path: path.resolve(__dirname, "../server/public"),
            filename: 'js/[name].js',
            publicPath: '/',
            chunkFilename: 'js/[id].js'
        },
        watch: true,
        plugins: [
            new webpack.DefinePlugin({
                'process.env': env
            }),
            new HtmlWebpackPlugin({
                template: "client/" + file + ".html",
                filename: path.resolve(__dirname, "../server/views/" + file + ".html"),
                root: path.resolve(__dirname, '../server/views'),
                minify: {
                    removeComments: false,
                    collapseWhitespace: false
                },
                inlineSource: '.(css)$'
            }),
            new CopyWebpackPlugin([{
                from: path.resolve(__dirname, '../client/static'),
                to: path.resolve(__dirname, "../server/public"),
                ignore: ['.*']
            }]),
            extractLess,
            new ExtractTextPlugin({
                filename: "css/[name].css",
                disable: false,
                allChunks: true
            })
        ],
        module: {
            rules: [{
                    test: /\.js$/,
                    exclude: /(node_modules|bower_components)/,
                    use: {
                        loader: 'babel-loader',
                        options: {
                            babelrc: true,
                            comments: false,
                            minified: false
                        }
                    }
                },
                {
                    test: /\.html$/,
                    use: [{
                        loader: 'html-loader'
                    }]
                },
                {
                    test: /\.css$/,
                    use: ExtractTextPlugin.extract({
                        fallback: "style-loader",
                        use: "css-loader"
                    })
                },
                {
                    test: /\.less$/,
                    use: extractLess.extract({
                        use: [{
                            loader: "css-loader"
                        }, {
                            loader: "less-loader"
                        }],
                        // use style-loader in development
                        fallback: "style-loader"
                    })
                },
                {
                    test: /\.(png|jpe?g|gif|svg)(\?.*)?$/,
                    use: [
                        'file-loader',
                        {
                            loader: 'file-loader',
                            options: {}
                        }
                    ],
                },
                {
                    test: /\.(mp4|webm|ogg|mp3|wav|flac|aac)(\?.*)?$/,
                    loader: 'url-loader',
                    options: {
                        limit: 10000,
                        name: "media/[name].[ext]"
                    }
                },
                {
                    test: /\.(woff2?|eot|ttf|otf)(\?.*)?$/,
                    loader: 'url-loader',
                    options: {
                        limit: 10000,
                        name: "fonts/[name].[ext]"
                    }
                }
            ]
        }
    }
    return webpackConfig;
}

module.exports = webpackConfigGenerator