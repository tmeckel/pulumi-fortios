// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure local FortiGuard Web Filter local ratings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.webfilter.Ftgdlocalrating("trname", {
 *     rating: "1",
 *     status: "enable",
 *     url: "sgala.com",
 * });
 * ```
 *
 * ## Import
 *
 * Webfilter FtgdLocalRating can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:webfilter/ftgdlocalrating:Ftgdlocalrating labelname {{url}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:webfilter/ftgdlocalrating:Ftgdlocalrating labelname {{url}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Ftgdlocalrating extends pulumi.CustomResource {
    /**
     * Get an existing Ftgdlocalrating resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FtgdlocalratingState, opts?: pulumi.CustomResourceOptions): Ftgdlocalrating {
        return new Ftgdlocalrating(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:webfilter/ftgdlocalrating:Ftgdlocalrating';

    /**
     * Returns true if the given object is an instance of Ftgdlocalrating.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ftgdlocalrating {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ftgdlocalrating.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Local rating.
     */
    public readonly rating!: pulumi.Output<string>;
    /**
     * Enable/disable local rating. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * URL to rate locally.
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Ftgdlocalrating resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FtgdlocalratingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FtgdlocalratingArgs | FtgdlocalratingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FtgdlocalratingState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["rating"] = state ? state.rating : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FtgdlocalratingArgs | undefined;
            if ((!args || args.rating === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rating'");
            }
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["rating"] = args ? args.rating : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ftgdlocalrating.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ftgdlocalrating resources.
 */
export interface FtgdlocalratingState {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Local rating.
     */
    rating?: pulumi.Input<string>;
    /**
     * Enable/disable local rating. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * URL to rate locally.
     */
    url?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ftgdlocalrating resource.
 */
export interface FtgdlocalratingArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Local rating.
     */
    rating: pulumi.Input<string>;
    /**
     * Enable/disable local rating. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * URL to rate locally.
     */
    url?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
