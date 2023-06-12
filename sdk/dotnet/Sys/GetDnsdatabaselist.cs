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
    public static class GetDnsdatabaselist
    {
        /// <summary>
        /// Provides a list of `fortios.sys.Dnsdatabase`.
        /// </summary>
        public static Task<GetDnsdatabaselistResult> InvokeAsync(GetDnsdatabaselistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDnsdatabaselistResult>("fortios:sys/getDnsdatabaselist:getDnsdatabaselist", args ?? new GetDnsdatabaselistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.sys.Dnsdatabase`.
        /// </summary>
        public static Output<GetDnsdatabaselistResult> Invoke(GetDnsdatabaselistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDnsdatabaselistResult>("fortios:sys/getDnsdatabaselist:getDnsdatabaselist", args ?? new GetDnsdatabaselistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDnsdatabaselistArgs : global::Pulumi.InvokeArgs
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

        public GetDnsdatabaselistArgs()
        {
        }
        public static new GetDnsdatabaselistArgs Empty => new GetDnsdatabaselistArgs();
    }

    public sealed class GetDnsdatabaselistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetDnsdatabaselistInvokeArgs()
        {
        }
        public static new GetDnsdatabaselistInvokeArgs Empty => new GetDnsdatabaselistInvokeArgs();
    }


    [OutputType]
    public sealed class GetDnsdatabaselistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.sys.Dnsdatabase`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetDnsdatabaselistResult(
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
