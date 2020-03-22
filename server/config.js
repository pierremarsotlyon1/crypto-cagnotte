module.exports = {
    name: 'API',
    env: process.env.NODE_ENV || 'development',
    port: process.env.PORT || 3000,
    base_url: process.env.BASE_URL || 'http://localhost:3000',
    db: {
        uri: process.env.MONGODB_URI || 'mongodb://127.0.0.1:27017/cagnotte',
    },
    jwt_key: {
        secret: "cagnotte_jwt_secret_pMFK7369IJKSQB198U247YAZHYDOZEIAGFB"
    }
};