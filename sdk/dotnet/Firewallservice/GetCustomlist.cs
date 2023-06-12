// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewallservice
{
    public static class GetCustomlist
    {
        /// <summary>
        /// Provides a list of `fortios.firewallservice.Custom`.
        /// </summary>
        public static Task<GetCustomlistResult> InvokeAsync(GetCustomlistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCustomlistResult>("fortios:firewallservice/getCustomlist:getCustomlist", args ?? new GetCustomlistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.firewallservice.Custom`.
        /// </summary>
        public static Output<GetCustomlistResult> Invoke(GetCustomlistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCustomlistResult>("fortios:firewallservice/getCustomlist:getCustomlist", args ?? new GetCustomlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCustomlistArgs : global::Pulumi.InvokeArgs
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

        public GetCustomlistArgs()
        {
        }
        public static new GetCustomlistArgs Empty => new GetCustomlistArgs();
    }

    public sealed class GetCustomlistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetCustomlistInvokeArgs()
        {
        }
        public static new GetCustomlistInvokeArgs Empty => new GetCustomlistInvokeArgs();
    }


    [OutputType]
    public sealed class GetCustomlistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.firewallservice.Custom`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetCustomlistResult(
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