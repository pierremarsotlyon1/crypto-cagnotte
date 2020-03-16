const Versement = require('./Versement');
const mongoose = require('mongoose');
const mongooseStringQuery = require('mongoose-string-query');
const timestamps = require('mongoose-timestamp');
const CagnotteSchema = new mongoose.Schema(
    {
        name: {
            type: String,
            required: true,
            trim: true,
        },
        walletAddress: {
            type: String,
            required: true,
            trim: true,
        },
        description: {
            type: String,
            required: true,
        },
        creator: {
            type: mongoose.ObjectId,
            required: true,
        },
        days: {
            type: Number,
            required: true,
        },
        versements: {
            type: [Versement],
        }
    },
    { minimize: false },
);
CagnotteSchema.plugin(timestamps);
CagnotteSchema.plugin(mongooseStringQuery);
const Cagnotte = mongoose.model('Cagnotte', CagnotteSchema);
module.exports = Cagnotte;