/**
 * Module Dependencies
 */
const errors = require('restify-errors');
const mongoose = require('mongoose');
import Coinbase from '../metier/Coinbase';

/**
 * Model Schema
 */
const Cagnotte = require('../models/Cagnotte');

export default class CagnotteController {
    coinbase = null;

    constructor(server) {
        server.get("/cagnottes", this.find);
        server.post("/cagnottes", this.add);
        server.get("/cagnottes/:id", this.get);
        server.del("/cagnottes/:id", this.delete);
        server.put("/cagnottes/:id", this.update);
        this.coinbase = new Coinbase();
    }

    find = (req, res, next) => {
        Cagnotte.find({}, {}, { sort: { createdAt: -1 } }, (err, cagnottes) => {
            if (err) {
                return next(new errors.InternalError("Erreur lors de la récupération des cagnottes"));
            }

            res.send(200, cagnottes);
            return next();
        });
    };

    add = (req, res, next) => {
        const cagnotte = new Cagnotte();
        cagnotte.name = req.body.name;
        cagnotte.description = req.body.description;
        cagnotte.days = req.body.days;
        cagnotte.creator = mongoose.Types.ObjectId(req.user.data)

        // Création de l'adresse USDC associée
        this.coinbase.getUSDCAddress()
            .then(walletAddress => {
                cagnotte.walletAddress = walletAddress;
                cagnotte.save(err => {
                    if (err) {
                        return next(new errors.InternalError("Erreur lors de la création de la cagnotte"));
                    }

                    res.send(200, cagnotte);
                    return next();
                });
            })
            .catch(error => {
                res.send(400, error);
                return next();
            });
    };

    get = (req, res, next) => {
        Cagnotte.findOne({ _id: mongoose.Types.ObjectId(req.params.id) }, (err, cagnotte) => {
            if (err) {
                console.log(err);
                return next(new errors.InternalError("Erreur lors de la récupération de la cagnotte"));
            }

            res.send(200, cagnotte);
            return next();
        });
    };

    delete = (req, res, next) => {
        Cagnotte.deleteOne({ _id: mongoose.Types.ObjectId(req.params.id) }, (err) => {
            if (err) {
                return next(new errors.InternalError("Erreur lors de la suppression de la cagnotte"));
            }

            res.send(200);
            return next();
        })
    };

    update = (req, res, next) => {
        next();
    };
}