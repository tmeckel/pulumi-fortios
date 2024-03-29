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

    public sealed class IsisIsisInterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication key-chain for level 1 PDUs.
        /// </summary>
        [Input("authKeychainL1")]
        public Input<string>? AuthKeychainL1 { get; set; }

        /// <summary>
        /// Authentication key-chain for level 2 PDUs.
        /// </summary>
        [Input("authKeychainL2")]
        public Input<string>? AuthKeychainL2 { get; set; }

        /// <summary>
        /// Level 1 authentication mode. Valid values: `md5`, `password`.
        /// </summary>
        [Input("authModeL1")]
        public Input<string>? AuthModeL1 { get; set; }

        /// <summary>
        /// Level 2 authentication mode. Valid values: `md5`, `password`.
        /// </summary>
        [Input("authModeL2")]
        public Input<string>? AuthModeL2 { get; set; }

        [Input("authPasswordL1")]
        private Input<string>? _authPasswordL1;

        /// <summary>
        /// Authentication password for level 1 PDUs.
        /// </summary>
        public Input<string>? AuthPasswordL1
        {
            get => _authPasswordL1;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _authPasswordL1 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("authPasswordL2")]
        private Input<string>? _authPasswordL2;

        /// <summary>
        /// Authentication password for level 2 PDUs.
        /// </summary>
        public Input<string>? AuthPasswordL2
        {
            get => _authPasswordL2;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _authPasswordL2 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Enable/disable authentication send-only for level 1 PDUs. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authSendOnlyL1")]
        public Input<string>? AuthSendOnlyL1 { get; set; }

        /// <summary>
        /// Enable/disable authentication send-only for level 2 PDUs. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("authSendOnlyL2")]
        public Input<string>? AuthSendOnlyL2 { get; set; }

        /// <summary>
        /// IS-IS interface's circuit type Valid values: `level-1-2`, `level-1`, `level-2`.
        /// </summary>
        [Input("circuitType")]
        public Input<string>? CircuitType { get; set; }

        /// <summary>
        /// Level 1 CSNP interval.
        /// </summary>
        [Input("csnpIntervalL1")]
        public Input<int>? CsnpIntervalL1 { get; set; }

        /// <summary>
        /// Level 2 CSNP interval.
        /// </summary>
        [Input("csnpIntervalL2")]
        public Input<int>? CsnpIntervalL2 { get; set; }

        /// <summary>
        /// Level 1 hello interval.
        /// </summary>
        [Input("helloIntervalL1")]
        public Input<int>? HelloIntervalL1 { get; set; }

        /// <summary>
        /// Level 2 hello interval.
        /// </summary>
        [Input("helloIntervalL2")]
        public Input<int>? HelloIntervalL2 { get; set; }

        /// <summary>
        /// Level 1 multiplier for Hello holding time.
        /// </summary>
        [Input("helloMultiplierL1")]
        public Input<int>? HelloMultiplierL1 { get; set; }

        /// <summary>
        /// Level 2 multiplier for Hello holding time.
        /// </summary>
        [Input("helloMultiplierL2")]
        public Input<int>? HelloMultiplierL2 { get; set; }

        /// <summary>
        /// Enable/disable padding to IS-IS hello packets. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("helloPadding")]
        public Input<string>? HelloPadding { get; set; }

        /// <summary>
        /// LSP transmission interval (milliseconds).
        /// </summary>
        [Input("lspInterval")]
        public Input<int>? LspInterval { get; set; }

        /// <summary>
        /// LSP retransmission interval (sec).
        /// </summary>
        [Input("lspRetransmitInterval")]
        public Input<int>? LspRetransmitInterval { get; set; }

        /// <summary>
        /// Enable/disable IS-IS mesh group. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("meshGroup")]
        public Input<string>? MeshGroup { get; set; }

        /// <summary>
        /// Mesh group ID &lt;0-4294967295&gt;, 0: mesh-group blocked.
        /// </summary>
        [Input("meshGroupId")]
        public Input<int>? MeshGroupId { get; set; }

        /// <summary>
        /// Level 1 metric for interface.
        /// </summary>
        [Input("metricL1")]
        public Input<int>? MetricL1 { get; set; }

        /// <summary>
        /// Level 2 metric for interface.
        /// </summary>
        [Input("metricL2")]
        public Input<int>? MetricL2 { get; set; }

        /// <summary>
        /// IS-IS interface name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// IS-IS interface's network type Valid values: `broadcast`, `point-to-point`, `loopback`.
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// Level 1 priority.
        /// </summary>
        [Input("priorityL1")]
        public Input<int>? PriorityL1 { get; set; }

        /// <summary>
        /// Level 2 priority.
        /// </summary>
        [Input("priorityL2")]
        public Input<int>? PriorityL2 { get; set; }

        /// <summary>
        /// Enable/disable interface for IS-IS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Enable/disable IPv6 interface for IS-IS. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status6")]
        public Input<string>? Status6 { get; set; }

        /// <summary>
        /// Level 1 wide metric for interface.
        /// </summary>
        [Input("wideMetricL1")]
        public Input<int>? WideMetricL1 { get; set; }

        /// <summary>
        /// Level 2 wide metric for interface.
        /// </summary>
        [Input("wideMetricL2")]
        public Input<int>? WideMetricL2 { get; set; }

        public IsisIsisInterfaceArgs()
        {
        }
        public static new IsisIsisInterfaceArgs Empty => new IsisIsisInterfaceArgs();
    }
}
