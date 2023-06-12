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

    public sealed class SdwanDuplicationArgs : global::Pulumi.ResourceArgs
    {
        [Input("dstaddr6s")]
        private InputList<Inputs.SdwanDuplicationDstaddr6Args>? _dstaddr6s;

        /// <summary>
        /// Destination address6 or address6 group names. The structure of `dstaddr6` block is documented below.
        /// </summary>
        public InputList<Inputs.SdwanDuplicationDstaddr6Args> Dstaddr6s
        {
            get => _dstaddr6s ?? (_dstaddr6s = new InputList<Inputs.SdwanDuplicationDstaddr6Args>());
            set => _dstaddr6s = value;
        }

        [Input("dstaddrs")]
        private InputList<Inputs.SdwanDuplicationDstaddrArgs>? _dstaddrs;

        /// <summary>
        /// Destination address or address group names. The structure of `dstaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.SdwanDuplicationDstaddrArgs> Dstaddrs
        {
            get => _dstaddrs ?? (_dstaddrs = new InputList<Inputs.SdwanDuplicationDstaddrArgs>());
            set => _dstaddrs = value;
        }

        [Input("dstintfs")]
        private InputList<Inputs.SdwanDuplicationDstintfArgs>? _dstintfs;

        /// <summary>
        /// Outgoing (egress) interfaces or zones. The structure of `dstintf` block is documented below.
        /// </summary>
        public InputList<Inputs.SdwanDuplicationDstintfArgs> Dstintfs
        {
            get => _dstintfs ?? (_dstintfs = new InputList<Inputs.SdwanDuplicationDstintfArgs>());
            set => _dstintfs = value;
        }

        /// <summary>
        /// Duplication rule ID (1 - 255).
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Enable/disable discarding of packets that have been duplicated. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("packetDeDuplication")]
        public Input<string>? PacketDeDuplication { get; set; }

        /// <summary>
        /// Configure packet duplication method. Valid values: `disable`, `force`, `on-demand`.
        /// </summary>
        [Input("packetDuplication")]
        public Input<string>? PacketDuplication { get; set; }

        [Input("serviceIds")]
        private InputList<Inputs.SdwanDuplicationServiceIdArgs>? _serviceIds;

        /// <summary>
        /// SD-WAN service rule ID list. The structure of `service_id` block is documented below.
        /// </summary>
        public InputList<Inputs.SdwanDuplicationServiceIdArgs> ServiceIds
        {
            get => _serviceIds ?? (_serviceIds = new InputList<Inputs.SdwanDuplicationServiceIdArgs>());
            set => _serviceIds = value;
        }

        [Input("services")]
        private InputList<Inputs.SdwanDuplicationServiceArgs>? _services;

        /// <summary>
        /// Service and service group name. The structure of `service` block is documented below.
        /// </summary>
        public InputList<Inputs.SdwanDuplicationServiceArgs> Services
        {
            get => _services ?? (_services = new InputList<Inputs.SdwanDuplicationServiceArgs>());
            set => _services = value;
        }

        /// <summary>
        /// Enable/disable packet duplication matching health-check SLAs in service rule. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("slaMatchService")]
        public Input<string>? SlaMatchService { get; set; }

        [Input("srcaddr6s")]
        private InputList<Inputs.SdwanDuplicationSrcaddr6Args>? _srcaddr6s;

        /// <summary>
        /// Source address6 or address6 group names. The structure of `srcaddr6` block is documented below.
        /// </summary>
        public InputList<Inputs.SdwanDuplicationSrcaddr6Args> Srcaddr6s
        {
            get => _srcaddr6s ?? (_srcaddr6s = new InputList<Inputs.SdwanDuplicationSrcaddr6Args>());
            set => _srcaddr6s = value;
        }

        [Input("srcaddrs")]
        private InputList<Inputs.SdwanDuplicationSrcaddrArgs>? _srcaddrs;

        /// <summary>
        /// Source address or address group names. The structure of `srcaddr` block is documented below.
        /// </summary>
        public InputList<Inputs.SdwanDuplicationSrcaddrArgs> Srcaddrs
        {
            get => _srcaddrs ?? (_srcaddrs = new InputList<Inputs.SdwanDuplicationSrcaddrArgs>());
            set => _srcaddrs = value;
        }

        [Input("srcintfs")]
        private InputList<Inputs.SdwanDuplicationSrcintfArgs>? _srcintfs;

        /// <summary>
        /// Incoming (ingress) interfaces or zones. The structure of `srcintf` block is documented below.
        /// </summary>
        public InputList<Inputs.SdwanDuplicationSrcintfArgs> Srcintfs
        {
            get => _srcintfs ?? (_srcintfs = new InputList<Inputs.SdwanDuplicationSrcintfArgs>());
            set => _srcintfs = value;
        }

        public SdwanDuplicationArgs()
        {
        }
        public static new SdwanDuplicationArgs Empty => new SdwanDuplicationArgs();
    }
}