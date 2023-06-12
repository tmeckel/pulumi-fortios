// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure debug URL addresses.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.webproxy.Debugurl("trname", {
 *     exact: "enable",
 *     status: "enable",
 *     urlPattern: "/examples/servlet/*Servlet",
 * });
 * ```
 *
 * ## Import
 *
 * WebProxy DebugUrl can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:webproxy/debugurl:Debugurl labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:webproxy/debugurl:Debugurl labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Debugurl extends pulumi.CustomResource {
    /**
     * Get an existing Debugurl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DebugurlState, opts?: pulumi.CustomResourceOptions): Debugurl {
        return new Debugurl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:webproxy/debugurl:Debugurl';

    /**
     * Returns true if the given object is an instance of Debugurl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Debugurl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Debugurl.__pulumiType;
    }

    /**
     * Enable/disable matching the exact path. Valid values: `enable`, `disable`.
     */
    public readonly exact!: pulumi.Output<string>;
    /**
     * Debug URL name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable this URL exemption. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * URL exemption pattern.
     */
    public readonly urlPattern!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Debugurl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DebugurlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DebugurlArgs | DebugurlState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DebugurlState | undefined;
            resourceInputs["exact"] = state ? state.exact : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["urlPattern"] = state ? state.urlPattern : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as DebugurlArgs | undefined;
            if ((!args || args.urlPattern === undefined) && !opts.urn) {
                throw new Error("Missing required property 'urlPattern'");
            }
            resourceInputs["exact"] = args ? args.exact : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["urlPattern"] = args ? args.urlPattern : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Debugurl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Debugurl resources.
 */
export interface DebugurlState {
    /**
     * Enable/disable matching the exact path. Valid values: `enable`, `disable`.
     */
    exact?: pulumi.Input<string>;
    /**
     * Debug URL name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable this URL exemption. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * URL exemption pattern.
     */
    urlPattern?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Debugurl resource.
 */
export interface DebugurlArgs {
    /**
     * Enable/disable matching the exact path. Valid values: `enable`, `disable`.
     */
    exact?: pulumi.Input<string>;
    /**
     * Debug URL name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable this URL exemption. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * URL exemption pattern.
     */
    urlPattern: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
