// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure access lists.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.router.Accesslist("trname", {comments: "test accesslist"});
 * ```
 * ## Note
 *
 * The feature can only be correctly supported when FortiOS Version >= 6.2.4, for FortiOS Version < 6.2.4, please use the following resource configuration as an alternative.
 *
 * ### Example
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname1 = new fortios.sys.Autoscript("trname1", {
 *     interval: 1,
 *     outputSize: 10,
 *     repeat: 1,
 *     script: `config router access-list
 * edit "static-redistribution"
 * config rule
 * edit 10
 * set prefix 10.0.0.0 255.255.255.0
 * set action permit
 * set exact-match enable
 * end
 * end
 *
 * `,
 *     start: "auto",
 * });
 * ```
 *
 * ## Import
 *
 * Router AccessList can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:router/accesslist:Accesslist labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:router/accesslist:Accesslist labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Accesslist extends pulumi.CustomResource {
    /**
     * Get an existing Accesslist resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccesslistState, opts?: pulumi.CustomResourceOptions): Accesslist {
        return new Accesslist(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:router/accesslist:Accesslist';

    /**
     * Returns true if the given object is an instance of Accesslist.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Accesslist {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Accesslist.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Rule. The structure of `rule` block is documented below.
     */
    public readonly rules!: pulumi.Output<outputs.router.AccesslistRule[] | undefined>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Accesslist resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AccesslistArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccesslistArgs | AccesslistState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccesslistState | undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as AccesslistArgs | undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Accesslist.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Accesslist resources.
 */
export interface AccesslistState {
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Rule. The structure of `rule` block is documented below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.router.AccesslistRule>[]>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Accesslist resource.
 */
export interface AccesslistArgs {
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Rule. The structure of `rule` block is documented below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.router.AccesslistRule>[]>;
    vdomparam?: pulumi.Input<string>;
}
