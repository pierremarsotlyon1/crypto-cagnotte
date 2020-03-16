const mongoose = require('mongoose');
const mongooseStringQuery = require('mongoose-string-query');
const timestamps = require('mongoose-timestamp');
const UserSchema = new mongoose.Schema(
    {
        email: {
            type: String,
            required: true,
            trim: true,
        },
        password: {
            type: String,
            required: true,
        },
        firstname: {
            type: String,
            required: true,
        },
        lastname: {
            type: String,
            required: true,
        }
    },
    { minimize: false },
);
UserSchema.plugin(timestamps);
UserSchema.plugin(mongooseStringQuery);
const User = mongoose.model('Users', UserSchema);
module.exports = User;