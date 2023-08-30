/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
import * as $protobuf from "protobufjs/minimal";

// Common aliases
const $util = $protobuf.util;

// Exported root namespace
const $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

export const http = $root.http = (() => {

    /**
     * Namespace http.
     * @exports http
     * @namespace
     */
    const http = {};

    http.RequestUsers = (function() {

        /**
         * Properties of a RequestUsers.
         * @memberof http
         * @interface IRequestUsers
         * @property {number|null} [page] RequestUsers page
         * @property {number|null} [page_size] RequestUsers page_size
         */

        /**
         * Constructs a new RequestUsers.
         * @memberof http
         * @classdesc Represents a RequestUsers.
         * @implements IRequestUsers
         * @constructor
         * @param {http.IRequestUsers=} [properties] Properties to set
         */
        function RequestUsers(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * RequestUsers page.
         * @member {number} page
         * @memberof http.RequestUsers
         * @instance
         */
        RequestUsers.prototype.page = 0;

        /**
         * RequestUsers page_size.
         * @member {number} page_size
         * @memberof http.RequestUsers
         * @instance
         */
        RequestUsers.prototype.page_size = 0;

        /**
         * Creates a new RequestUsers instance using the specified properties.
         * @function create
         * @memberof http.RequestUsers
         * @static
         * @param {http.IRequestUsers=} [properties] Properties to set
         * @returns {http.RequestUsers} RequestUsers instance
         */
        RequestUsers.create = function create(properties) {
            return new RequestUsers(properties);
        };

        return RequestUsers;
    })();

    http.User = (function() {

        /**
         * Properties of a User.
         * @memberof http
         * @interface IUser
         * @property {number|null} [id] User id
         * @property {string|null} [username] User username
         * @property {string|null} [email] User email
         */

        /**
         * Constructs a new User.
         * @memberof http
         * @classdesc Represents a User.
         * @implements IUser
         * @constructor
         * @param {http.IUser=} [properties] Properties to set
         */
        function User(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * User id.
         * @member {number} id
         * @memberof http.User
         * @instance
         */
        User.prototype.id = 0;

        /**
         * User username.
         * @member {string} username
         * @memberof http.User
         * @instance
         */
        User.prototype.username = "";

        /**
         * User email.
         * @member {string} email
         * @memberof http.User
         * @instance
         */
        User.prototype.email = "";

        /**
         * Creates a new User instance using the specified properties.
         * @function create
         * @memberof http.User
         * @static
         * @param {http.IUser=} [properties] Properties to set
         * @returns {http.User} User instance
         */
        User.create = function create(properties) {
            return new User(properties);
        };

        return User;
    })();

    http.ResponseUsers = (function() {

        /**
         * Properties of a ResponseUsers.
         * @memberof http
         * @interface IResponseUsers
         * @property {Array.<http.IUser>|null} [data_list] ResponseUsers data_list
         */

        /**
         * Constructs a new ResponseUsers.
         * @memberof http
         * @classdesc Represents a ResponseUsers.
         * @implements IResponseUsers
         * @constructor
         * @param {http.IResponseUsers=} [properties] Properties to set
         */
        function ResponseUsers(properties) {
            this.data_list = [];
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * ResponseUsers data_list.
         * @member {Array.<http.IUser>} data_list
         * @memberof http.ResponseUsers
         * @instance
         */
        ResponseUsers.prototype.data_list = $util.emptyArray;

        /**
         * Creates a new ResponseUsers instance using the specified properties.
         * @function create
         * @memberof http.ResponseUsers
         * @static
         * @param {http.IResponseUsers=} [properties] Properties to set
         * @returns {http.ResponseUsers} ResponseUsers instance
         */
        ResponseUsers.create = function create(properties) {
            return new ResponseUsers(properties);
        };

        return ResponseUsers;
    })();

    http.RequestNil = (function() {

        /**
         * Properties of a RequestNil.
         * @memberof http
         * @interface IRequestNil
         */

        /**
         * Constructs a new RequestNil.
         * @memberof http
         * @classdesc Represents a RequestNil.
         * @implements IRequestNil
         * @constructor
         * @param {http.IRequestNil=} [properties] Properties to set
         */
        function RequestNil(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new RequestNil instance using the specified properties.
         * @function create
         * @memberof http.RequestNil
         * @static
         * @param {http.IRequestNil=} [properties] Properties to set
         * @returns {http.RequestNil} RequestNil instance
         */
        RequestNil.create = function create(properties) {
            return new RequestNil(properties);
        };

        return RequestNil;
    })();

    http.ResponseNil = (function() {

        /**
         * Properties of a ResponseNil.
         * @memberof http
         * @interface IResponseNil
         */

        /**
         * Constructs a new ResponseNil.
         * @memberof http
         * @classdesc Represents a ResponseNil.
         * @implements IResponseNil
         * @constructor
         * @param {http.IResponseNil=} [properties] Properties to set
         */
        function ResponseNil(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new ResponseNil instance using the specified properties.
         * @function create
         * @memberof http.ResponseNil
         * @static
         * @param {http.IResponseNil=} [properties] Properties to set
         * @returns {http.ResponseNil} ResponseNil instance
         */
        ResponseNil.create = function create(properties) {
            return new ResponseNil(properties);
        };

        return ResponseNil;
    })();

    http.RequestPostEmpty = (function() {

        /**
         * Properties of a RequestPostEmpty.
         * @memberof http
         * @interface IRequestPostEmpty
         * @property {string|null} [name] RequestPostEmpty name
         */

        /**
         * Constructs a new RequestPostEmpty.
         * @memberof http
         * @classdesc Represents a RequestPostEmpty.
         * @implements IRequestPostEmpty
         * @constructor
         * @param {http.IRequestPostEmpty=} [properties] Properties to set
         */
        function RequestPostEmpty(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * RequestPostEmpty name.
         * @member {string} name
         * @memberof http.RequestPostEmpty
         * @instance
         */
        RequestPostEmpty.prototype.name = "";

        /**
         * Creates a new RequestPostEmpty instance using the specified properties.
         * @function create
         * @memberof http.RequestPostEmpty
         * @static
         * @param {http.IRequestPostEmpty=} [properties] Properties to set
         * @returns {http.RequestPostEmpty} RequestPostEmpty instance
         */
        RequestPostEmpty.create = function create(properties) {
            return new RequestPostEmpty(properties);
        };

        return RequestPostEmpty;
    })();

    http.CommonNil = (function() {

        /**
         * Properties of a CommonNil.
         * @memberof http
         * @interface ICommonNil
         */

        /**
         * Constructs a new CommonNil.
         * @memberof http
         * @classdesc Represents a CommonNil.
         * @implements ICommonNil
         * @constructor
         * @param {http.ICommonNil=} [properties] Properties to set
         */
        function CommonNil(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Creates a new CommonNil instance using the specified properties.
         * @function create
         * @memberof http.CommonNil
         * @static
         * @param {http.ICommonNil=} [properties] Properties to set
         * @returns {http.CommonNil} CommonNil instance
         */
        CommonNil.create = function create(properties) {
            return new CommonNil(properties);
        };

        return CommonNil;
    })();

    return http;
})();

export const google = $root.google = (() => {

    /**
     * Namespace google.
     * @exports google
     * @namespace
     */
    const google = {};

    google.protobuf = (function() {

        /**
         * Namespace protobuf.
         * @memberof google
         * @namespace
         */
        const protobuf = {};

        protobuf.Empty = (function() {

            /**
             * Properties of an Empty.
             * @memberof google.protobuf
             * @interface IEmpty
             */

            /**
             * Constructs a new Empty.
             * @memberof google.protobuf
             * @classdesc Represents an Empty.
             * @implements IEmpty
             * @constructor
             * @param {google.protobuf.IEmpty=} [properties] Properties to set
             */
            function Empty(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * Creates a new Empty instance using the specified properties.
             * @function create
             * @memberof google.protobuf.Empty
             * @static
             * @param {google.protobuf.IEmpty=} [properties] Properties to set
             * @returns {google.protobuf.Empty} Empty instance
             */
            Empty.create = function create(properties) {
                return new Empty(properties);
            };

            return Empty;
        })();

        return protobuf;
    })();

    return google;
})();

export { $root as default };
