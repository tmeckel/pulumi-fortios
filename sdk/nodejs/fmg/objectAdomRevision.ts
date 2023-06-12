// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource supports Create/Read/Update/Delete object adom revision for FortiManager.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test1 = new fortios.fmg.ObjectAdomRevision("test1", {
 *     createdBy: "fortinet",
 *     description: "adom revision",
 *     locked: 0,
 * });
 * ```
 */
export class ObjectAdomRevision extends pulumi.CustomResource {
    /**
     * Get an existing ObjectAdomRevision resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ObjectAdomRevisionState, opts?: pulumi.CustomResourceOptions): ObjectAdomRevision {
        return new ObjectAdomRevision(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:fmg/objectAdomRevision:ObjectAdomRevision';

    /**
     * Returns true if the given object is an instance of ObjectAdomRevision.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ObjectAdomRevision {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ObjectAdomRevision.__pulumiType;
    }

    /**
     * ADOM name. default is 'root'.
     */
    public readonly adom!: pulumi.Output<string | undefined>;
    /**
     * Who created this adom revision.
     */
    public readonly createdBy!: pulumi.Output<string | undefined>;
    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * lock. 0 means unlock and 1 means locked.
     */
    public readonly locked!: pulumi.Output<number | undefined>;
    /**
     * Adom revision name.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a ObjectAdomRevision resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ObjectAdomRevisionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ObjectAdomRevisionArgs | ObjectAdomRevisionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ObjectAdomRevisionState | undefined;
            resourceInputs["adom"] = state ? state.adom : undefined;
            resourceInputs["createdBy"] = state ? state.createdBy : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["locked"] = state ? state.locked : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ObjectAdomRevisionArgs | undefined;
            resourceInputs["adom"] = args ? args.adom : undefined;
            resourceInputs["createdBy"] = args ? args.createdBy : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["locked"] = args ? args.locked : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ObjectAdomRevision.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ObjectAdomRevision resources.
 */
export interface ObjectAdomRevisionState {
    /**
     * ADOM name. default is 'root'.
     */
    adom?: pulumi.Input<string>;
    /**
     * Who created this adom revision.
     */
    createdBy?: pulumi.Input<string>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * lock. 0 means unlock and 1 means locked.
     */
    locked?: pulumi.Input<number>;
    /**
     * Adom revision name.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ObjectAdomRevision resource.
 */
export interface ObjectAdomRevisionArgs {
    /**
     * ADOM name. default is 'root'.
     */
    adom?: pulumi.Input<string>;
    /**
     * Who created this adom revision.
     */
    createdBy?: pulumi.Input<string>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * lock. 0 means unlock and 1 means locked.
     */
    locked?: pulumi.Input<number>;
    /**
     * Adom revision name.
     */
    name?: pulumi.Input<string>;
}
