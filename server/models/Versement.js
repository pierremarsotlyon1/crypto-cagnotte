const mongoose = require('mongoose');
const mongooseStringQuery = require('mongoose-string-query');
const timestamps = require('mongoose-timestamp');
const VersementSchema = new mongoose.Schema(
    {
        cagnotte: {
            type: mongoose.ObjectId,
            required: true,
        },
        total: {
            type: Number,
            required: true,
        }
    },
    { minimize: false },
);
VersementSchema.plugin(timestamps);
VersementSchema.plugin(mongooseStringQuery);
const Versement = mongoose.model('Versement', VersementSchema);
module.exports = Versement;