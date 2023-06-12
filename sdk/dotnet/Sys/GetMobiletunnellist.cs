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
    public static class GetMobiletunnellist
    {
        /// <summary>
        /// Provides a list of `fortios.sys.Mobiletunnel`.
        /// </summary>
        public static Task<GetMobiletunnellistResult> InvokeAsync(GetMobiletunnellistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMobiletunnellistResult>("fortios:sys/getMobiletunnellist:getMobiletunnellist", args ?? new GetMobiletunnellistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.sys.Mobiletunnel`.
        /// </summary>
        public static Output<GetMobiletunnellistResult> Invoke(GetMobiletunnellistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMobiletunnellistResult>("fortios:sys/getMobiletunnellist:getMobiletunnellist", args ?? new GetMobiletunnellistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMobiletunnellistArgs : global::Pulumi.InvokeArgs
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

        public GetMobiletunnellistArgs()
        {
        }
        public static new GetMobiletunnellistArgs Empty => new GetMobiletunnellistArgs();
    }

    public sealed class GetMobiletunnellistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetMobiletunnellistInvokeArgs()
        {
        }
        public static new GetMobiletunnellistInvokeArgs Empty => new GetMobiletunnellistInvokeArgs();
    }


    [OutputType]
    public sealed class GetMobiletunnellistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.sys.Mobiletunnel`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetMobiletunnellistResult(
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
