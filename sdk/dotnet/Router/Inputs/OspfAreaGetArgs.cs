// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Inputs
{

    public sealed class OspfAreaGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication type.
        /// </summary>
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comments")]
        public Input<string>? Comments { get; set; }

        /// <summary>
        /// Summary default cost of stub or NSSA area.
        /// </summary>
        [Input("defaultCost")]
        public Input<int>? DefaultCost { get; set; }

        [Input("filterLists")]
        private InputList<Inputs.OspfAreaFilterListGetArgs>? _filterLists;

        /// <summary>
        /// OSPF area filter-list configuration. The structure of `filter_list` block is documented below.
        /// </summary>
        public InputList<Inputs.OspfAreaFilterListGetArgs> FilterLists
        {
            get => _filterLists ?? (_filterLists = new InputList<Inputs.OspfAreaFilterListGetArgs>());
            set => _filterLists = value;
        }

        /// <summary>
        /// Area entry IP address.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Redistribute, advertise, or do not originate Type-7 default route into NSSA area. Valid values: `enable`, `always`, `disable`.
        /// </summary>
        [Input("nssaDefaultInformationOriginate")]
        public Input<string>? NssaDefaultInformationOriginate { get; set; }

        /// <summary>
        /// OSPF default metric.
        /// </summary>
        [Input("nssaDefaultInformationOriginateMetric")]
        public Input<int>? NssaDefaultInformationOriginateMetric { get; set; }

        /// <summary>
        /// OSPF metric type for default routes. Valid values: `1`, `2`.
        /// </summary>
        [Input("nssaDefaultInformationOriginateMetricType")]
        public Input<string>? NssaDefaultInformationOriginateMetricType { get; set; }

        /// <summary>
        /// Enable/disable redistribute into NSSA area. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nssaRedistribution")]
        public Input<string>? NssaRedistribution { get; set; }

        /// <summary>
        /// NSSA translator role type. Valid values: `candidate`, `never`, `always`.
        /// </summary>
        [Input("nssaTranslatorRole")]
        public Input<string>? NssaTranslatorRole { get; set; }

        [Input("ranges")]
        private InputList<Inputs.OspfAreaRangeGetArgs>? _ranges;

        /// <summary>
        /// OSPF area range configuration. The structure of `range` block is documented below.
        /// </summary>
        public InputList<Inputs.OspfAreaRangeGetArgs> Ranges
        {
            get => _ranges ?? (_ranges = new InputList<Inputs.OspfAreaRangeGetArgs>());
            set => _ranges = value;
        }

        /// <summary>
        /// Enable/disable shortcut option. Valid values: `disable`, `enable`, `default`.
        /// </summary>
        [Input("shortcut")]
        public Input<string>? Shortcut { get; set; }

        /// <summary>
        /// Stub summary setting. Valid values: `no-summary`, `summary`.
        /// </summary>
        [Input("stubType")]
        public Input<string>? StubType { get; set; }

        /// <summary>
        /// Area type setting. Valid values: `regular`, `nssa`, `stub`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("virtualLinks")]
        private InputList<Inputs.OspfAreaVirtualLinkGetArgs>? _virtualLinks;

        /// <summary>
        /// OSPF virtual link configuration. The structure of `virtual_link` block is documented below.
        /// </summary>
        public InputList<Inputs.OspfAreaVirtualLinkGetArgs> VirtualLinks
        {
            get => _virtualLinks ?? (_virtualLinks = new InputList<Inputs.OspfAreaVirtualLinkGetArgs>());
            set => _virtualLinks = value;
        }

        public OspfAreaGetArgs()
        {
        }
        public static new OspfAreaGetArgs Empty => new OspfAreaGetArgs();
    }
}
