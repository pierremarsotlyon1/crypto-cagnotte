/**
 * Module Dependencies
 */
const errors = require('restify-errors');
const bcrypt = require("bcryptjs")
const SALT = 10;
const JWT_SECRET = "cagnotte_jwt_secret_pMFK7369IJKSQB198U247YAZHYDOZEIAGFB";
const jwt = require("jsonwebtoken");
/**
 * Model Schema
 */
const User = require('../models/User');
module.exports = function (server) {

    /**
     * Login
     */
    server.post("/login", (req, res, next) => {
        // On récupère l'utilisateur via son email
        User.findOne({ email: req.body.email }, (err, doc) => {
            if (err) {
                return next(new errors.InternalError(err.message));
            }

            // Check de l'user
            if (!doc) {
                return next(new errors.InternalError("L'utilisateur n'existe pas"));
            }

            // On vérifie son mot de passe
            bcrypt.compare(req.body.password, doc.password, (err, isMatch) => {
                if (err || !isMatch) {
                    return next(new errors.InternalError("Le mot de passe ne correspond pas"));
                }

                const token = jwt.sign({ data: doc._id }, JWT_SECRET);
                res.send(200, token);
                next();
            })
        })
    });

    /**
     * Création de compte
     */
    server.post("/register", (req, res, next) => {
        const body = req.body;

        if (body.password !== body.confirmPassword) {
            return next(new errors.InternalError("Les mots de passe ne correspondent pas"));
        }

        // On récupère l'utilisateur via son email
        User.findOne({ email: req.body.email }, (err, doc) => {
            if (err) {
                return next(new errors.InternalError(err.message));
            }

            // Check de l'user
            if (doc) {
                return next(new errors.InternalError("L'email est déjà pris"));
            }

            const user = new User();
            user.email = body.email;
            user.firstname = body.firstname;
            user.lastname = body.lastname;

            bcrypt.hash(body.password, SALT, function (err, hash) {
                user.password = hash;
                user.save((err) => {
                    if (err) {
                        return next(new errors.InternalError(err.message));
                    }

                    const token = jwt.sign({ data: user._id }, JWT_SECRET);
                    res.send(201, token);
                    next();
                });
            });
        })
    });
};