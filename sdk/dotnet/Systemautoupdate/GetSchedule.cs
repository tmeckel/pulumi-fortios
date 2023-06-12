// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Systemautoupdate
{
    public static class GetSchedule
    {
        /// <summary>
        /// Use this data source to get information on fortios systemautoupdate schedule
        /// </summary>
        public static Task<GetScheduleResult> InvokeAsync(GetScheduleArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetScheduleResult>("fortios:systemautoupdate/getSchedule:getSchedule", args ?? new GetScheduleArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios systemautoupdate schedule
        /// </summary>
        public static Output<GetScheduleResult> Invoke(GetScheduleInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetScheduleResult>("fortios:systemautoupdate/getSchedule:getSchedule", args ?? new GetScheduleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetScheduleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetScheduleArgs()
        {
        }
        public static new GetScheduleArgs Empty => new GetScheduleArgs();
    }

    public sealed class GetScheduleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetScheduleInvokeArgs()
        {
        }
        public static new GetScheduleInvokeArgs Empty => new GetScheduleInvokeArgs();
    }


    [OutputType]
    public sealed class GetScheduleResult
    {
        /// <summary>
        /// Update day.
        /// </summary>
        public readonly string Day;
        /// <summary>
        /// Update frequency.
        /// </summary>
        public readonly string Frequency;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Enable/disable scheduled updates.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Update time.
        /// </summary>
        public readonly string Time;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetScheduleResult(
            string day,

            string frequency,

            string id,

            string status,

            string time,

            string? vdomparam)
        {
            Day = day;
            Frequency = frequency;
            Id = id;
            Status = status;
            Time = time;
            Vdomparam = vdomparam;
        }
    }
}