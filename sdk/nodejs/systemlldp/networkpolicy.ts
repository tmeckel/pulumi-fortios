// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure LLDP network policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.systemlldp.Networkpolicy("trname", {comment: "test"});
 * ```
 *
 * ## Import
 *
 * SystemLldp NetworkPolicy can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:systemlldp/networkpolicy:Networkpolicy labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:systemlldp/networkpolicy:Networkpolicy labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Networkpolicy extends pulumi.CustomResource {
    /**
     * Get an existing Networkpolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkpolicyState, opts?: pulumi.CustomResourceOptions): Networkpolicy {
        return new Networkpolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:systemlldp/networkpolicy:Networkpolicy';

    /**
     * Returns true if the given object is an instance of Networkpolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Networkpolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Networkpolicy.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Guest. The structure of `guest` block is documented below.
     */
    public readonly guest!: pulumi.Output<outputs.systemlldp.NetworkpolicyGuest>;
    /**
     * Guest Voice Signaling. The structure of `guestVoiceSignaling` block is documented below.
     */
    public readonly guestVoiceSignaling!: pulumi.Output<outputs.systemlldp.NetworkpolicyGuestVoiceSignaling>;
    /**
     * LLDP network policy name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Softphone. The structure of `softphone` block is documented below.
     */
    public readonly softphone!: pulumi.Output<outputs.systemlldp.NetworkpolicySoftphone>;
    /**
     * Streaming Video. The structure of `streamingVideo` block is documented below.
     */
    public readonly streamingVideo!: pulumi.Output<outputs.systemlldp.NetworkpolicyStreamingVideo>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Video Conferencing. The structure of `videoConferencing` block is documented below.
     */
    public readonly videoConferencing!: pulumi.Output<outputs.systemlldp.NetworkpolicyVideoConferencing>;
    /**
     * Video Signaling. The structure of `videoSignaling` block is documented below.
     */
    public readonly videoSignaling!: pulumi.Output<outputs.systemlldp.NetworkpolicyVideoSignaling>;
    /**
     * Voice. The structure of `voice` block is documented below.
     */
    public readonly voice!: pulumi.Output<outputs.systemlldp.NetworkpolicyVoice>;
    /**
     * Voice signaling. The structure of `voiceSignaling` block is documented below.
     */
    public readonly voiceSignaling!: pulumi.Output<outputs.systemlldp.NetworkpolicyVoiceSignaling>;

    /**
     * Create a Networkpolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NetworkpolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkpolicyArgs | NetworkpolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkpolicyState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["guest"] = state ? state.guest : undefined;
            resourceInputs["guestVoiceSignaling"] = state ? state.guestVoiceSignaling : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["softphone"] = state ? state.softphone : undefined;
            resourceInputs["streamingVideo"] = state ? state.streamingVideo : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["videoConferencing"] = state ? state.videoConferencing : undefined;
            resourceInputs["videoSignaling"] = state ? state.videoSignaling : undefined;
            resourceInputs["voice"] = state ? state.voice : undefined;
            resourceInputs["voiceSignaling"] = state ? state.voiceSignaling : undefined;
        } else {
            const args = argsOrState as NetworkpolicyArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["guest"] = args ? args.guest : undefined;
            resourceInputs["guestVoiceSignaling"] = args ? args.guestVoiceSignaling : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["softphone"] = args ? args.softphone : undefined;
            resourceInputs["streamingVideo"] = args ? args.streamingVideo : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["videoConferencing"] = args ? args.videoConferencing : undefined;
            resourceInputs["videoSignaling"] = args ? args.videoSignaling : undefined;
            resourceInputs["voice"] = args ? args.voice : undefined;
            resourceInputs["voiceSignaling"] = args ? args.voiceSignaling : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Networkpolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Networkpolicy resources.
 */
export interface NetworkpolicyState {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Guest. The structure of `guest` block is documented below.
     */
    guest?: pulumi.Input<inputs.systemlldp.NetworkpolicyGuest>;
    /**
     * Guest Voice Signaling. The structure of `guestVoiceSignaling` block is documented below.
     */
    guestVoiceSignaling?: pulumi.Input<inputs.systemlldp.NetworkpolicyGuestVoiceSignaling>;
    /**
     * LLDP network policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Softphone. The structure of `softphone` block is documented below.
     */
    softphone?: pulumi.Input<inputs.systemlldp.NetworkpolicySoftphone>;
    /**
     * Streaming Video. The structure of `streamingVideo` block is documented below.
     */
    streamingVideo?: pulumi.Input<inputs.systemlldp.NetworkpolicyStreamingVideo>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Video Conferencing. The structure of `videoConferencing` block is documented below.
     */
    videoConferencing?: pulumi.Input<inputs.systemlldp.NetworkpolicyVideoConferencing>;
    /**
     * Video Signaling. The structure of `videoSignaling` block is documented below.
     */
    videoSignaling?: pulumi.Input<inputs.systemlldp.NetworkpolicyVideoSignaling>;
    /**
     * Voice. The structure of `voice` block is documented below.
     */
    voice?: pulumi.Input<inputs.systemlldp.NetworkpolicyVoice>;
    /**
     * Voice signaling. The structure of `voiceSignaling` block is documented below.
     */
    voiceSignaling?: pulumi.Input<inputs.systemlldp.NetworkpolicyVoiceSignaling>;
}

/**
 * The set of arguments for constructing a Networkpolicy resource.
 */
export interface NetworkpolicyArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Guest. The structure of `guest` block is documented below.
     */
    guest?: pulumi.Input<inputs.systemlldp.NetworkpolicyGuest>;
    /**
     * Guest Voice Signaling. The structure of `guestVoiceSignaling` block is documented below.
     */
    guestVoiceSignaling?: pulumi.Input<inputs.systemlldp.NetworkpolicyGuestVoiceSignaling>;
    /**
     * LLDP network policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Softphone. The structure of `softphone` block is documented below.
     */
    softphone?: pulumi.Input<inputs.systemlldp.NetworkpolicySoftphone>;
    /**
     * Streaming Video. The structure of `streamingVideo` block is documented below.
     */
    streamingVideo?: pulumi.Input<inputs.systemlldp.NetworkpolicyStreamingVideo>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Video Conferencing. The structure of `videoConferencing` block is documented below.
     */
    videoConferencing?: pulumi.Input<inputs.systemlldp.NetworkpolicyVideoConferencing>;
    /**
     * Video Signaling. The structure of `videoSignaling` block is documented below.
     */
    videoSignaling?: pulumi.Input<inputs.systemlldp.NetworkpolicyVideoSignaling>;
    /**
     * Voice. The structure of `voice` block is documented below.
     */
    voice?: pulumi.Input<inputs.systemlldp.NetworkpolicyVoice>;
    /**
     * Voice signaling. The structure of `voiceSignaling` block is documented below.
     */
    voiceSignaling?: pulumi.Input<inputs.systemlldp.NetworkpolicyVoiceSignaling>;
}
