// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Inputs
{

    public sealed class AdminGuiDashboardWidgetGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fabric device to monitor.
        /// </summary>
        [Input("fabricDevice")]
        public Input<string>? FabricDevice { get; set; }

        [Input("filters")]
        private InputList<Inputs.AdminGuiDashboardWidgetFilterGetArgs>? _filters;

        /// <summary>
        /// FortiView filters. The structure of `filters` block is documented below.
        /// </summary>
        public InputList<Inputs.AdminGuiDashboardWidgetFilterGetArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.AdminGuiDashboardWidgetFilterGetArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Height.
        /// </summary>
        [Input("height")]
        public Input<int>? Height { get; set; }

        /// <summary>
        /// Widget ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Security Audit Rating industry. Valid values: `default`, `custom`.
        /// </summary>
        [Input("industry")]
        public Input<string>? Industry { get; set; }

        /// <summary>
        /// Interface to monitor.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Security Audit Rating region. Valid values: `default`, `custom`.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Field to aggregate the data by. Valid values: `source`, `destination`, `country`, `intfpair`, `srcintf`, `dstintf`, `policy`, `wificlient`, `shaper`, `endpoint-vulnerability`, `endpoint-device`, `application`, `cloud-app`, `cloud-user`, `web-domain`, `web-category`, `web-search-phrase`, `threat`, `system`, `unauth`, `admin`, `vpn`.
        /// </summary>
        [Input("reportBy")]
        public Input<string>? ReportBy { get; set; }

        /// <summary>
        /// Field to sort the data by.
        /// </summary>
        [Input("sortBy")]
        public Input<string>? SortBy { get; set; }

        /// <summary>
        /// Timeframe period of reported data. Valid values: `realtime`, `5min`, `hour`, `day`, `week`.
        /// </summary>
        [Input("timeframe")]
        public Input<string>? Timeframe { get; set; }

        /// <summary>
        /// Widget title.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// Widget type. Valid values: `sysinfo`, `licinfo`, `vminfo`, `forticloud`, `cpu-usage`, `memory-usage`, `disk-usage`, `log-rate`, `sessions`, `session-rate`, `tr-history`, `analytics`, `usb-modem`, `admins`, `security-fabric`, `security-fabric-ranking`, `ha-status`, `vulnerability-summary`, `host-scan-summary`, `fortiview`, `botnet-activity`, `fortimail`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Visualization to use. Valid values: `table`, `bubble`, `country`, `chord`.
        /// </summary>
        [Input("visualization")]
        public Input<string>? Visualization { get; set; }

        /// <summary>
        /// Width.
        /// </summary>
        [Input("width")]
        public Input<int>? Width { get; set; }

        /// <summary>
        /// X position.
        /// </summary>
        [Input("xPos")]
        public Input<int>? XPos { get; set; }

        /// <summary>
        /// Y position.
        /// </summary>
        [Input("yPos")]
        public Input<int>? YPos { get; set; }

        public AdminGuiDashboardWidgetGetArgs()
        {
        }
        public static new AdminGuiDashboardWidgetGetArgs Empty => new AdminGuiDashboardWidgetGetArgs();
    }
}
