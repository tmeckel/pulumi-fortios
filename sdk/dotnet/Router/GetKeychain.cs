// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router
{
    public static class GetKeychain
    {
        /// <summary>
        /// Use this data source to get information on an fortios router keychain
        /// </summary>
        public static Task<GetKeychainResult> InvokeAsync(GetKeychainArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKeychainResult>("fortios:router/getKeychain:getKeychain", args ?? new GetKeychainArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios router keychain
        /// </summary>
        public static Output<GetKeychainResult> Invoke(GetKeychainInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKeychainResult>("fortios:router/getKeychain:getKeychain", args ?? new GetKeychainInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKeychainArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired router keychain.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetKeychainArgs()
        {
        }
        public static new GetKeychainArgs Empty => new GetKeychainArgs();
    }

    public sealed class GetKeychainInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired router keychain.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetKeychainInvokeArgs()
        {
        }
        public static new GetKeychainInvokeArgs Empty => new GetKeychainInvokeArgs();
    }


    [OutputType]
    public sealed class GetKeychainResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Configuration method to edit key settings. The structure of `key` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKeychainKeyResult> Keys;
        /// <summary>
        /// Key-chain name.
        /// </summary>
        public readonly string Name;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetKeychainResult(
            string id,

            ImmutableArray<Outputs.GetKeychainKeyResult> keys,

            string name,

            string? vdomparam)
        {
            Id = id;
            Keys = keys;
            Name = name;
            Vdomparam = vdomparam;
        }
    }
}
