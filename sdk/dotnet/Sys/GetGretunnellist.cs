// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys
{
    public static class GetGretunnellist
    {
        /// <summary>
        /// Provides a list of `fortios.sys.Gretunnel`.
        /// </summary>
        public static Task<GetGretunnellistResult> InvokeAsync(GetGretunnellistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGretunnellistResult>("fortios:sys/getGretunnellist:getGretunnellist", args ?? new GetGretunnellistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.sys.Gretunnel`.
        /// </summary>
        public static Output<GetGretunnellistResult> Invoke(GetGretunnellistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGretunnellistResult>("fortios:sys/getGretunnellist:getGretunnellist", args ?? new GetGretunnellistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGretunnellistArgs : global::Pulumi.InvokeArgs
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

        public GetGretunnellistArgs()
        {
        }
        public static new GetGretunnellistArgs Empty => new GetGretunnellistArgs();
    }

    public sealed class GetGretunnellistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetGretunnellistInvokeArgs()
        {
        }
        public static new GetGretunnellistInvokeArgs Empty => new GetGretunnellistInvokeArgs();
    }


    [OutputType]
    public sealed class GetGretunnellistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.sys.Gretunnel`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetGretunnellistResult(
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
