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
    public static class GetAutomationactionlist
    {
        /// <summary>
        /// Provides a list of `fortios.sys.Automationaction`.
        /// </summary>
        public static Task<GetAutomationactionlistResult> InvokeAsync(GetAutomationactionlistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAutomationactionlistResult>("fortios:sys/getAutomationactionlist:getAutomationactionlist", args ?? new GetAutomationactionlistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.sys.Automationaction`.
        /// </summary>
        public static Output<GetAutomationactionlistResult> Invoke(GetAutomationactionlistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAutomationactionlistResult>("fortios:sys/getAutomationactionlist:getAutomationactionlist", args ?? new GetAutomationactionlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAutomationactionlistArgs : global::Pulumi.InvokeArgs
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

        public GetAutomationactionlistArgs()
        {
        }
        public static new GetAutomationactionlistArgs Empty => new GetAutomationactionlistArgs();
    }

    public sealed class GetAutomationactionlistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetAutomationactionlistInvokeArgs()
        {
        }
        public static new GetAutomationactionlistInvokeArgs Empty => new GetAutomationactionlistInvokeArgs();
    }


    [OutputType]
    public sealed class GetAutomationactionlistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.sys.Automationaction`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetAutomationactionlistResult(
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
