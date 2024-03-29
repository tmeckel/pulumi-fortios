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
    public static class GetReplacemsggroup
    {
        /// <summary>
        /// Use this data source to get information on an fortios system replacemsggroup
        /// </summary>
        public static Task<GetReplacemsggroupResult> InvokeAsync(GetReplacemsggroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetReplacemsggroupResult>("fortios:sys/getReplacemsggroup:getReplacemsggroup", args ?? new GetReplacemsggroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system replacemsggroup
        /// </summary>
        public static Output<GetReplacemsggroupResult> Invoke(GetReplacemsggroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetReplacemsggroupResult>("fortios:sys/getReplacemsggroup:getReplacemsggroup", args ?? new GetReplacemsggroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReplacemsggroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system replacemsggroup.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetReplacemsggroupArgs()
        {
        }
        public static new GetReplacemsggroupArgs Empty => new GetReplacemsggroupArgs();
    }

    public sealed class GetReplacemsggroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system replacemsggroup.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetReplacemsggroupInvokeArgs()
        {
        }
        public static new GetReplacemsggroupInvokeArgs Empty => new GetReplacemsggroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetReplacemsggroupResult
    {
        /// <summary>
        /// Replacement message table entries. The structure of `admin` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplacemsggroupAdminResult> Admins;
        /// <summary>
        /// Replacement message table entries. The structure of `alertmail` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplacemsggroupAlertmailResult> Alertmails;
        /// <summary>
        /// Replacement message table entries. The structure of `auth` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplacemsggroupAuthResult> Auths;
        /// <summary>
        /// Replacement message table entries. The structure of `automation` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplacemsggroupAutomationResult> Automations;
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comment;
        /// <summary>
        /// Replacement message table entries. The structure of `custom_message` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplacemsggroupCustomMessageResult> CustomMessages;
        /// <summary>
        /// Replacement message table entries. The structure of `device_detection_portal` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplacemsggroupDeviceDetectionPortalResult> DeviceDetectionPortals;
        /// <summary>
        /// Replacement message table entries. The structure of `ec` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplacemsggroupEcResult> Ecs;
        /// <summary>
        /// Replacement message table entries. The structure of `fortiguard_wf` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplacemsggroupFortiguardWfResult> FortiguardWfs;
        /// <summary>
        /// Replacement message table entries. The structure of `ftp` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplacemsggroupFtpResult> Ftps;
        /// <summary>
        /// Group type.
        /// </summary>
        public readonly string GroupType;
        /// <summary>
        /// Replacement message table entries. The structure of `http` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplacemsggroupHttpResult> Https;
        /// <summary>
        /// Replacement message table entries. The structure of `icap` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplacemsggroupIcapResult> Icaps;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Replacement message table entries. The structure of `mail` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplacemsggroupMailResult> Mails;
        /// <summary>
        /// Replacement message table entries. The structure of `nac_quar` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplacemsggroupNacQuarResult> NacQuars;
        /// <summary>
        /// Group name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Replacement message table entries. The structure of `nntp` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplacemsggroupNntpResult> Nntps;
        /// <summary>
        /// Replacement message table entries. The structure of `spam` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplacemsggroupSpamResult> Spams;
        /// <summary>
        /// Replacement message table entries. The structure of `sslvpn` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplacemsggroupSslvpnResult> Sslvpns;
        /// <summary>
        /// Replacement message table entries. The structure of `traffic_quota` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplacemsggroupTrafficQuotaResult> TrafficQuotas;
        /// <summary>
        /// Replacement message table entries. The structure of `utm` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplacemsggroupUtmResult> Utms;
        public readonly string? Vdomparam;
        /// <summary>
        /// Replacement message table entries. The structure of `webproxy` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReplacemsggroupWebproxyResult> Webproxies;

        [OutputConstructor]
        private GetReplacemsggroupResult(
            ImmutableArray<Outputs.GetReplacemsggroupAdminResult> admins,

            ImmutableArray<Outputs.GetReplacemsggroupAlertmailResult> alertmails,

            ImmutableArray<Outputs.GetReplacemsggroupAuthResult> auths,

            ImmutableArray<Outputs.GetReplacemsggroupAutomationResult> automations,

            string comment,

            ImmutableArray<Outputs.GetReplacemsggroupCustomMessageResult> customMessages,

            ImmutableArray<Outputs.GetReplacemsggroupDeviceDetectionPortalResult> deviceDetectionPortals,

            ImmutableArray<Outputs.GetReplacemsggroupEcResult> ecs,

            ImmutableArray<Outputs.GetReplacemsggroupFortiguardWfResult> fortiguardWfs,

            ImmutableArray<Outputs.GetReplacemsggroupFtpResult> ftps,

            string groupType,

            ImmutableArray<Outputs.GetReplacemsggroupHttpResult> https,

            ImmutableArray<Outputs.GetReplacemsggroupIcapResult> icaps,

            string id,

            ImmutableArray<Outputs.GetReplacemsggroupMailResult> mails,

            ImmutableArray<Outputs.GetReplacemsggroupNacQuarResult> nacQuars,

            string name,

            ImmutableArray<Outputs.GetReplacemsggroupNntpResult> nntps,

            ImmutableArray<Outputs.GetReplacemsggroupSpamResult> spams,

            ImmutableArray<Outputs.GetReplacemsggroupSslvpnResult> sslvpns,

            ImmutableArray<Outputs.GetReplacemsggroupTrafficQuotaResult> trafficQuotas,

            ImmutableArray<Outputs.GetReplacemsggroupUtmResult> utms,

            string? vdomparam,

            ImmutableArray<Outputs.GetReplacemsggroupWebproxyResult> webproxies)
        {
            Admins = admins;
            Alertmails = alertmails;
            Auths = auths;
            Automations = automations;
            Comment = comment;
            CustomMessages = customMessages;
            DeviceDetectionPortals = deviceDetectionPortals;
            Ecs = ecs;
            FortiguardWfs = fortiguardWfs;
            Ftps = ftps;
            GroupType = groupType;
            Https = https;
            Icaps = icaps;
            Id = id;
            Mails = mails;
            NacQuars = nacQuars;
            Name = name;
            Nntps = nntps;
            Spams = spams;
            Sslvpns = sslvpns;
            TrafficQuotas = trafficQuotas;
            Utms = utms;
            Vdomparam = vdomparam;
            Webproxies = webproxies;
        }
    }
}
