// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class SecurityPolicyseq extends pulumi.CustomResource {
    /**
     * Get an existing SecurityPolicyseq resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityPolicyseqState, opts?: pulumi.CustomResourceOptions): SecurityPolicyseq {
        return new SecurityPolicyseq(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/securityPolicyseq:SecurityPolicyseq';

    /**
     * Returns true if the given object is an instance of SecurityPolicyseq.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityPolicyseq {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityPolicyseq.__pulumiType;
    }

    /**
     * The alter position: should only be "before" or "after"
     */
    public readonly alterPosition!: pulumi.Output<string>;
    /**
     * Comment
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Enable status detection for policySrcId and policy_dst_id
     */
    public readonly enableStateChecking!: pulumi.Output<boolean | undefined>;
    /**
     * The dest policy id which you want to alter
     */
    public readonly policyDstId!: pulumi.Output<number>;
    /**
     * The policy id which you want to alter
     */
    public readonly policySrcId!: pulumi.Output<number>;
    public /*out*/ readonly statePolicyLists!: pulumi.Output<outputs.firewall.SecurityPolicyseqStatePolicyList[]>;
    public readonly statePolicySrcdstPos!: pulumi.Output<string | undefined>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SecurityPolicyseq resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityPolicyseqArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityPolicyseqArgs | SecurityPolicyseqState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityPolicyseqState | undefined;
            resourceInputs["alterPosition"] = state ? state.alterPosition : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["enableStateChecking"] = state ? state.enableStateChecking : undefined;
            resourceInputs["policyDstId"] = state ? state.policyDstId : undefined;
            resourceInputs["policySrcId"] = state ? state.policySrcId : undefined;
            resourceInputs["statePolicyLists"] = state ? state.statePolicyLists : undefined;
            resourceInputs["statePolicySrcdstPos"] = state ? state.statePolicySrcdstPos : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SecurityPolicyseqArgs | undefined;
            if ((!args || args.alterPosition === undefined) && !opts.urn) {
                throw new Error("Missing required property 'alterPosition'");
            }
            if ((!args || args.policyDstId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyDstId'");
            }
            if ((!args || args.policySrcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policySrcId'");
            }
            resourceInputs["alterPosition"] = args ? args.alterPosition : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["enableStateChecking"] = args ? args.enableStateChecking : undefined;
            resourceInputs["policyDstId"] = args ? args.policyDstId : undefined;
            resourceInputs["policySrcId"] = args ? args.policySrcId : undefined;
            resourceInputs["statePolicySrcdstPos"] = args ? args.statePolicySrcdstPos : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["statePolicyLists"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityPolicyseq.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityPolicyseq resources.
 */
export interface SecurityPolicyseqState {
    /**
     * The alter position: should only be "before" or "after"
     */
    alterPosition?: pulumi.Input<string>;
    /**
     * Comment
     */
    comment?: pulumi.Input<string>;
    /**
     * Enable status detection for policySrcId and policy_dst_id
     */
    enableStateChecking?: pulumi.Input<boolean>;
    /**
     * The dest policy id which you want to alter
     */
    policyDstId?: pulumi.Input<number>;
    /**
     * The policy id which you want to alter
     */
    policySrcId?: pulumi.Input<number>;
    statePolicyLists?: pulumi.Input<pulumi.Input<inputs.firewall.SecurityPolicyseqStatePolicyList>[]>;
    statePolicySrcdstPos?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityPolicyseq resource.
 */
export interface SecurityPolicyseqArgs {
    /**
     * The alter position: should only be "before" or "after"
     */
    alterPosition: pulumi.Input<string>;
    /**
     * Comment
     */
    comment?: pulumi.Input<string>;
    /**
     * Enable status detection for policySrcId and policy_dst_id
     */
    enableStateChecking?: pulumi.Input<boolean>;
    /**
     * The dest policy id which you want to alter
     */
    policyDstId: pulumi.Input<number>;
    /**
     * The policy id which you want to alter
     */
    policySrcId: pulumi.Input<number>;
    statePolicySrcdstPos?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
