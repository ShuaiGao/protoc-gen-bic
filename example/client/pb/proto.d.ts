import * as $protobuf from "protobufjs";
import Long = require("long");
/** Namespace http. */
export namespace http {

    /** Properties of a RequestUsers. */
    interface IRequestUsers {

        /** RequestUsers page */
        page?: (number|null);

        /** RequestUsers page_size */
        page_size?: (number|null);
    }

    /** Represents a RequestUsers. */
    class RequestUsers implements IRequestUsers {

        /**
         * Constructs a new RequestUsers.
         * @param [properties] Properties to set
         */
        constructor(properties?: http.IRequestUsers);

        /** RequestUsers page. */
        public page: number;

        /** RequestUsers page_size. */
        public page_size: number;

        /**
         * Creates a new RequestUsers instance using the specified properties.
         * @param [properties] Properties to set
         * @returns RequestUsers instance
         */
        public static create(properties?: http.IRequestUsers): http.RequestUsers;
    }

    /** Properties of a User. */
    interface IUser {

        /** User id */
        id?: (number|null);

        /** User username */
        username?: (string|null);

        /** User email */
        email?: (string|null);
    }

    /** Represents a User. */
    class User implements IUser {

        /**
         * Constructs a new User.
         * @param [properties] Properties to set
         */
        constructor(properties?: http.IUser);

        /** User id. */
        public id: number;

        /** User username. */
        public username: string;

        /** User email. */
        public email: string;

        /**
         * Creates a new User instance using the specified properties.
         * @param [properties] Properties to set
         * @returns User instance
         */
        public static create(properties?: http.IUser): http.User;
    }

    /** Properties of a ResponseUsers. */
    interface IResponseUsers {

        /** ResponseUsers data_list */
        data_list?: (http.IUser[]|null);
    }

    /** Represents a ResponseUsers. */
    class ResponseUsers implements IResponseUsers {

        /**
         * Constructs a new ResponseUsers.
         * @param [properties] Properties to set
         */
        constructor(properties?: http.IResponseUsers);

        /** ResponseUsers data_list. */
        public data_list: http.IUser[];

        /**
         * Creates a new ResponseUsers instance using the specified properties.
         * @param [properties] Properties to set
         * @returns ResponseUsers instance
         */
        public static create(properties?: http.IResponseUsers): http.ResponseUsers;
    }

    /** Properties of a RequestNil. */
    interface IRequestNil {
    }

    /** Represents a RequestNil. */
    class RequestNil implements IRequestNil {

        /**
         * Constructs a new RequestNil.
         * @param [properties] Properties to set
         */
        constructor(properties?: http.IRequestNil);

        /**
         * Creates a new RequestNil instance using the specified properties.
         * @param [properties] Properties to set
         * @returns RequestNil instance
         */
        public static create(properties?: http.IRequestNil): http.RequestNil;
    }

    /** Properties of a ResponseNil. */
    interface IResponseNil {
    }

    /** Represents a ResponseNil. */
    class ResponseNil implements IResponseNil {

        /**
         * Constructs a new ResponseNil.
         * @param [properties] Properties to set
         */
        constructor(properties?: http.IResponseNil);

        /**
         * Creates a new ResponseNil instance using the specified properties.
         * @param [properties] Properties to set
         * @returns ResponseNil instance
         */
        public static create(properties?: http.IResponseNil): http.ResponseNil;
    }

    /** Properties of a RequestPostEmpty. */
    interface IRequestPostEmpty {

        /** RequestPostEmpty name */
        name?: (string|null);
    }

    /** Represents a RequestPostEmpty. */
    class RequestPostEmpty implements IRequestPostEmpty {

        /**
         * Constructs a new RequestPostEmpty.
         * @param [properties] Properties to set
         */
        constructor(properties?: http.IRequestPostEmpty);

        /** RequestPostEmpty name. */
        public name: string;

        /**
         * Creates a new RequestPostEmpty instance using the specified properties.
         * @param [properties] Properties to set
         * @returns RequestPostEmpty instance
         */
        public static create(properties?: http.IRequestPostEmpty): http.RequestPostEmpty;
    }

    /** Properties of a CommonNil. */
    interface ICommonNil {
    }

    /** Represents a CommonNil. */
    class CommonNil implements ICommonNil {

        /**
         * Constructs a new CommonNil.
         * @param [properties] Properties to set
         */
        constructor(properties?: http.ICommonNil);

        /**
         * Creates a new CommonNil instance using the specified properties.
         * @param [properties] Properties to set
         * @returns CommonNil instance
         */
        public static create(properties?: http.ICommonNil): http.CommonNil;
    }
}

/** Namespace google. */
export namespace google {

    /** Namespace protobuf. */
    namespace protobuf {

        /** Properties of an Empty. */
        interface IEmpty {
        }

        /** Represents an Empty. */
        class Empty implements IEmpty {

            /**
             * Constructs a new Empty.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IEmpty);

            /**
             * Creates a new Empty instance using the specified properties.
             * @param [properties] Properties to set
             * @returns Empty instance
             */
            public static create(properties?: google.protobuf.IEmpty): google.protobuf.Empty;
        }
    }
}
