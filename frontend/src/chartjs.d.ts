import { Tooltip } from "chart.js";

declare module 'chart.js' {
    interface TooltipPositionerMap {
      CustomPositioner: TooltipPositionerFunction<ChartType>;
    }
  }