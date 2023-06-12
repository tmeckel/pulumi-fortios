// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys.Outputs
{

    [OutputType]
    public sealed class GetAdminGuiDashboardWidgetResult
    {
        /// <summary>
        /// Fabric device to monitor.
        /// </summary>
        public readonly string FabricDevice;
        /// <summary>
        /// FortiView filters. The structure of `filters` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAdminGuiDashboardWidgetFilterResult> Filters;
        /// <summary>
        /// Height.
        /// </summary>
        public readonly int Height;
        /// <summary>
        /// Select menu ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Security Audit Rating industry.
        /// </summary>
        public readonly string Industry;
        /// <summary>
        /// Interface to monitor.
        /// </summary>
        public readonly string Interface;
        /// <summary>
        /// Security Audit Rating region.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Field to aggregate the data by.
        /// </summary>
        public readonly string ReportBy;
        /// <summary>
        /// Field to sort the data by.
        /// </summary>
        public readonly string SortBy;
        /// <summary>
        /// Timeframe period of reported data.
        /// </summary>
        public readonly string Timeframe;
        /// <summary>
        /// Widget title.
        /// </summary>
        public readonly string Title;
        /// <summary>
        /// Widget type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Visualization to use.
        /// </summary>
        public readonly string Visualization;
        /// <summary>
        /// Width.
        /// </summary>
        public readonly int Width;
        /// <summary>
        /// X position.
        /// </summary>
        public readonly int XPos;
        /// <summary>
        /// Y position.
        /// </summary>
        public readonly int YPos;

        [OutputConstructor]
        private GetAdminGuiDashboardWidgetResult(
            string fabricDevice,

            ImmutableArray<Outputs.GetAdminGuiDashboardWidgetFilterResult> filters,

            int height,

            int id,

            string industry,

            string @interface,

            string region,

            string reportBy,

            string sortBy,

            string timeframe,

            string title,

            string type,

            string visualization,

            int width,

            int xPos,

            int yPos)
        {
            FabricDevice = fabricDevice;
            Filters = filters;
            Height = height;
            Id = id;
            Industry = industry;
            Interface = @interface;
            Region = region;
            ReportBy = reportBy;
            SortBy = sortBy;
            Timeframe = timeframe;
            Title = title;
            Type = type;
            Visualization = visualization;
            Width = width;
            XPos = xPos;
            YPos = yPos;
        }
    }
}
