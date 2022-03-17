// const TsconfigPathsPlugin = require('tsconfig-paths-webpack-plugin');

module.exports = {
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: ['source-map-loader', "babel-loader"],
            }, {
                test: /\.css$/i,
                use: ["style-loader", "css-loader"],
            }
        ]
    },
    resolve: {
        extensions: ['.js', '.jsx'],
        plugins: [],
        fallback: {
            fs: false
        }
    },
    performance: {
        hints: false,
        maxEntrypointSize: 512000,
        maxAssetSize: 512000
    }
};