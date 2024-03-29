// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios system replacemsgimage
 */
export function getReplacemsgimage(args: GetReplacemsgimageArgs, opts?: pulumi.InvokeOptions): Promise<GetReplacemsgimageResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:sys/getReplacemsgimage:getReplacemsgimage", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getReplacemsgimage.
 */
export interface GetReplacemsgimageArgs {
    /**
     * Specify the name of the desired system replacemsgimage.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getReplacemsgimage.
 */
export interface GetReplacemsgimageResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Image data.
     */
    readonly imageBase64: string;
    /**
     * Image type.
     */
    readonly imageType: string;
    /**
     * Image name.
     */
    readonly name: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios system replacemsgimage
 */
export function getReplacemsgimageOutput(args: GetReplacemsgimageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReplacemsgimageResult> {
    return pulumi.output(args).apply((a: any) => getReplacemsgimage(a, opts))
}

/**
 * A collection of arguments for invoking getReplacemsgimage.
 */
export interface GetReplacemsgimageOutputArgs {
    /**
     * Specify the name of the desired system replacemsgimage.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
