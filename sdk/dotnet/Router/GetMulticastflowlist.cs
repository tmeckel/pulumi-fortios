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
    public static class GetMulticastflowlist
    {
        /// <summary>
        /// Provides a list of `fortios.router.Multicastflow`.
        /// </summary>
        public static Task<GetMulticastflowlistResult> InvokeAsync(GetMulticastflowlistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMulticastflowlistResult>("fortios:router/getMulticastflowlist:getMulticastflowlist", args ?? new GetMulticastflowlistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.router.Multicastflow`.
        /// </summary>
        public static Output<GetMulticastflowlistResult> Invoke(GetMulticastflowlistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMulticastflowlistResult>("fortios:router/getMulticastflowlist:getMulticastflowlist", args ?? new GetMulticastflowlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMulticastflowlistArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter used to scope the list. See Filter results of datasource.
        /// </summary>
        [Input("filter")]
        public string? Filter { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetMulticastflowlistArgs()
        {
        }
        public static new GetMulticastflowlistArgs Empty => new GetMulticastflowlistArgs();
    }

    public sealed class GetMulticastflowlistInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter used to scope the list. See Filter results of datasource.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetMulticastflowlistInvokeArgs()
        {
        }
        public static new GetMulticastflowlistInvokeArgs Empty => new GetMulticastflowlistInvokeArgs();
    }


    [OutputType]
    public sealed class GetMulticastflowlistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.router.Multicastflow`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetMulticastflowlistResult(
            string? filter,

            string id,

            ImmutableArray<string> namelists,

            string? vdomparam)
        {
            Filter = filter;
            Id = id;
            Namelists = namelists;
            Vdomparam = vdomparam;
        }
    }
}
