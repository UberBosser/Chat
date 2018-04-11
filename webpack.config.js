const webpack = require("webpack");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");


module.exports = {
    entry: {
        index: "./src/containers/index.jsx",
        404: "./src/containers/404.jsx"
    },

    output: {
        path: __dirname + "/static/js/",
        filename: "[name].bundle.js"
    },

    module: {
        rules: [
            {
                resolve: {
                    alias: {
                        inferno: __dirname + "/node_modules/inferno/dist/index.dev.esm.js" 
                    } 
                }
            },
            {            
                test: /\.(js|jsx)$/,
                loader: "babel-loader",
                exclude: /node_modules/
            },
            {
                test: /\.(css|sass)$/,
                use: [
                    MiniCssExtractPlugin.loader,
                    {
                        loader: "css-loader" 
                    },
                    {
                        loader: "sass-loader" 
                    }
                ],
                exclude: /node_modules/
            },
            {
                test: /\.(png|jp(e*)g|svg)$/,
                use: [{
                    loader: "url-loader",
                    options: {
                        limit: 8192,
                        name: "[name].[hash].img.[ext]",
                        outputPath: "../images/",
                        publicPath: "/static/images/"
                    }
                    }]
            }
        ]
    },

    plugins: [
        new webpack.DefinePlugin({
            'process.env': {
                'NODE_ENV': JSON.stringify('production')
            }
        }),
        new MiniCssExtractPlugin({
            filename: "../css/[name].bundle.css",
        })
    ],

    mode: "development"
};
