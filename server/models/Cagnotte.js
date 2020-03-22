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
            ref:'Users'
        },
        days: {
            type: Number,
            required: true,
        },
        versements: {
            type: [Versement.schema],
        }
    },
    { minimize: false },
);
CagnotteSchema.plugin(timestamps);
CagnotteSchema.plugin(mongooseStringQuery);
const Cagnotte = mongoose.model('Cagnottes', CagnotteSchema);
module.exports = Cagnotte;