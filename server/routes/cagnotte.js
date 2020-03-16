/**
 * Module Dependencies
 */
const errors = require('restify-errors');

/**
 * Model Schema
 */
const Cagnotte = require('../models/Cagnotte');
module.exports = function (server) {

    /**
     * Permet de récupérer la liste des cagnottes
     */
    server.get("/cagnottes", (req, res, next) => {
        next();
    });

    /**
     * Permet de créer une cagnotte
     */
    server.post("/cagnottes", (req, res, next) => {
        next();
    });

    /**
     * Permet de récupérer une cagnotte
     */
    server.get("/cagnottes/:id", (req, res, next) => {
        next();
    });

    /**
     * Permet de supprimer une cagnotte
     */
    server.delete("/cagnottes/:id", (req, res, next) => {
        next();
    });

    /**
     * Permet de modifier une cagnotte
     */
    server.put("/cagnottes/:id", (req, res, next) => {
        next();
    });
};