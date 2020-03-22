const axios = require('axios');

const API_ENDPOINT = "https://api.coinbase.com/";
const API_COINBASE_VERSION = "v2";
const API_KEY = "5NfXuQ60G4otrGQj";
const API_SECRET = "2sC5yLeObVry1TZfYsEdkwTPYZvX9G4U";
const crypto = require('crypto');

const USDCAccount = "5be9d0134f06a1033d4073fb";

export default class Coinbase {
    getUSDCAddress = () => {
        const pathUrl = "/accounts/" + USDCAccount + "/addresses";
        const body = {};
        const headers = this.getHeaders(body, pathUrl, "POST");

        return axios.post(API_ENDPOINT + API_COINBASE_VERSION + pathUrl, body, {
            headers
        })
            .then(response => {
                return response.data.data.address;
            })
            .catch(e => {
                return null;
            });
    };

    getHeaders = (body, url, method) => {
        const timestamp = Math.floor(Date.now() / 1000);
        const bodyStr = body ? JSON.stringify(body) : '';
        const message = timestamp + method + ("/" + API_COINBASE_VERSION + url) + bodyStr;
        const signature = crypto.createHmac('sha256', API_SECRET).update(message).digest("hex");

        return {
            'CB-ACCESS-SIGN': signature,
            'CB-ACCESS-TIMESTAMP': timestamp,
            'CB-ACCESS-KEY': API_KEY,
            'Content-Type': 'application/json',
        };
    };
}