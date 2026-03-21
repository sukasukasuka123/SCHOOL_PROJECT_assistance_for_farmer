# -*- coding: utf-8 -*-
"""
果园守护者 · 四会专项
模块2：现实落地盈收仿真（3年期）

核心设计思路：
  不追求"技术最优"，而是追求"钱算得过"。
  三种情景的关键差异不在定价，而在付费转化率和政府合作进度——
  这才是大创落地阶段真正的不确定性来源。

运行方法：python sim_revenue.py
输出文件：sim_revenue.png（保存在脚本同目录下）
"""

import numpy as np
import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
import matplotlib.font_manager as fm
from matplotlib.gridspec import GridSpec
import os
import warnings
warnings.filterwarnings('ignore')

# ── 字体设置（Windows / Mac / Linux 自动适配）──────────────────
def get_chinese_font():
    candidates = [
        'Microsoft YaHei',    # Windows 微软雅黑
        'SimHei',             # Windows 黑体
        'PingFang SC',        # macOS
        'Noto Sans CJK SC',   # Linux
        'WenQuanYi Micro Hei',
        'DejaVu Sans',        # 最终兜底
    ]
    available = {f.name for f in fm.fontManager.ttflist}
    for font in candidates:
        if font in available:
            return font
    return None

font_name = get_chinese_font()
if font_name:
    plt.rcParams['font.family'] = font_name
plt.rcParams.update({
    'axes.unicode_minus': False,
    'figure.dpi':         150,
    'axes.titlesize':     12,
    'axes.labelsize':     10,
    'xtick.labelsize':    9,
    'ytick.labelsize':    9,
    'legend.fontsize':    9,
    'axes.spines.top':    False,
    'axes.spines.right':  False,
})

RED    = '#c0392b'
AMBER  = '#d4870a'
FOREST = '#2d5a27'
TEAL   = '#0f766e'
SLATE  = '#374151'
LGRAY  = '#f3f4f6'
MGRAY  = '#9ca3af'

# ══════════════════════════════════════════════════════════════════
# 成本结构（按大创团队轻资产打法估算）
# ══════════════════════════════════════════════════════════════════

# 一次性前期投入（CAPEX）
CAPEX_ITEMS = {
    '侦察机自制×2架（机架+飞控+传感器）': 8_000,
    '投放机自制×1架（含投放模块）':        12_000,
    '边缘计算模块（树莓派+摄像头+LoRa）': 3_000,
    'App/小程序前后端对接外包':            5_000,
}
CAPEX = sum(CAPEX_ITEMS.values())   # 合计 2.8万元

# 月度固定运营成本
MONTHLY_FIXED = {
    '云服务器+联盟链节点租用':   500,    # 元/月
    '无人机维护耗材（电池/零件）': 600,  # 元/月，约设备价值 1.7%/月
    '系统折旧分摊':               CAPEX / 36,  # 3年线性折旧
}
MONTHLY_FIXED_TOTAL = sum(MONTHLY_FIXED.values())   # 约 1,878元/月

# 月度变动成本（随客户数线性增长）
COST_PER_CLIENT = 80   # 元/客户/月（人工售后+驻点巡检补贴）

# ══════════════════════════════════════════════════════════════════
# 三种落地情景参数
# ══════════════════════════════════════════════════════════════════
#
# 设计逻辑：
#   保守情景 → 农民不信任、推广受阻、没有政府采购
#   基准情景 → 示范点成功、口碑传播、植保站合作落地
#   乐观情景 → 政策顺风、媒体曝光、溯源渠道打通
#
# 最关键的变量：付费转化率（接触了系统的农民中愿意付钱的比例）
# 其次是：政府采购是否落地、溯源溢价分润是否实现

SCENARIOS = {
    '保守情景\n（推广受阻）': {
        'color': RED,
        'ls':    '--',
        # 每月新增"接触"客户数（包括免费示范、口碑介绍等）
        # 前6个月只有免费示范点，几乎没有商业线索
        'monthly_leads': [0]*6 + [0.6]*6 + [1.0]*12 + [1.2]*12,

        # 付费转化率：接触客户中愿意签约付费的比例
        # 保守情景：农民观望心态重，四个接触的只有一个付钱
        'conversion_rate': 0.25,

        # 平均签约面积（亩）：保守情景以小散户为主
        'avg_acreage': 55,

        # 服务单价（元/亩/季）：低于定价，因为没有议价能力
        'price_per_mu': 68,

        # 每客户年均服务季数（沙糖桔一年有春秋两季嫩梢期）
        'seasons_per_year': 1.5,

        # 年流失率（客户觉得效果不稳定或自己放弃种桔）
        'annual_churn': 0.18,

        # 政府/植保站采购年收入（元）× 落地概率
        # 保守情景：没有拿到政府订单
        'gov_contract_prob': 0.0,
        'gov_contract_value': 150_000,   # 如果拿到约15万/年

        # 溯源溢价分润：需要打通盒马/叮咚等渠道，大创阶段很难
        'premium_prob': 0.0,
        'premium_per_mu_year': 270,      # 溯源溢价平台分润约270元/亩/年

        # 政府采购开始月份（第几个月才有可能谈成）
        'gov_start_month': 99,  # 保守情景：谈不成，设为极大值
    },

    '基准情景\n（正常推进）': {
        'color': FOREST,
        'ls':    '-',
        'monthly_leads':     [0]*3 + [0.8]*5 + [1.5]*8 + [2.2]*12 + [2.8]*8,
        'conversion_rate':   0.40,
        'avg_acreage':       82,
        'price_per_mu':      78,
        'seasons_per_year':  1.8,
        'annual_churn':      0.08,
        'gov_contract_prob': 0.40,   # 40%概率拿到植保站合同
        'gov_contract_value':150_000,
        'premium_prob':      0.25,   # 25%概率打通一个渠道
        'premium_per_mu_year':270,
        'gov_start_month':   9,      # 第9个月开始有可能谈成
    },

    '乐观情景\n（政策顺风）': {
        'color': TEAL,
        'ls':    '-.',
        'monthly_leads':     [0]*2 + [1.2]*3 + [2.5]*5 + [4.0]*8 + [5.0]*10 + [6.0]*8,
        'conversion_rate':   0.55,
        'avg_acreage':       108,
        'price_per_mu':      85,
        'seasons_per_year':  2.0,
        'annual_churn':      0.05,
        'gov_contract_prob': 0.70,
        'gov_contract_value':200_000,   # 更大的采购合同
        'premium_prob':      0.55,
        'premium_per_mu_year':270,
        'gov_start_month':   7,
    },
}

MONTHS = 36   # 仿真周期：3年

# ══════════════════════════════════════════════════════════════════
# 盈收仿真主函数
# ══════════════════════════════════════════════════════════════════
def run_revenue_sim(sc_name):
    """
    逐月模拟一个情景下的收支情况

    客户增长模型：
      - 每月产生若干"接触客户"（免费示范、口碑推荐）
      - 接触后有3个月决策周期，之后按转化率付费
      - 每月有流失（合同到期不续或果农放弃种桔）

    收入结构：
      1. 服务订阅（主力）= 客户数 × 亩数 × 单价 × 季数/12
      2. 政府采购（关键增量）= 按月均摊，有概率落地
      3. 溯源分润（远期收益）= 需渠道成熟，大创阶段不依赖

    成本结构：
      = 固定成本（折旧+云服务器+维护）+ 变动成本（售后人力）
    """
    p = SCENARIOS[sc_name]
    leads = p['monthly_leads']
    assert len(leads) == MONTHS, f"monthly_leads 长度应为 {MONTHS}"

    active_clients = 0.0    # 当月活跃付费客户数
    cum_profit     = -CAPEX # 累计净利润（初始已投入CAPEX）

    monthly_revenue  = []
    monthly_cost     = []
    monthly_clients  = []
    cumulative_profit = []
    monthly_profit   = []

    # 政府合同是否落地（用随机数模拟概率，固定seed保证可重复）
    rng = np.random.default_rng(seed=123)
    gov_landed = rng.random() < p['gov_contract_prob']

    for m in range(MONTHS):

        # ── 新增付费客户（3个月决策周期）──────────────────────────
        if m >= 3:
            # 3个月前接触的客户，现在按转化率完成付费签约
            newly_paid = leads[m - 3] * p['conversion_rate']
            active_clients += newly_paid

        # ── 客户流失（年流失率转月）────────────────────────────────
        monthly_churn_rate = p['annual_churn'] / 12
        active_clients = max(0.0, active_clients * (1 - monthly_churn_rate))

        # ── 收入计算 ────────────────────────────────────────────────

        # 1. 服务订阅收入（主力收入）
        sub_revenue = (active_clients
                       * p['avg_acreage']
                       * p['price_per_mu']
                       * p['seasons_per_year']
                       / 12)

        # 2. 政府采购收入（需达到谈判月份且合同成功落地）
        gov_revenue = 0.0
        if gov_landed and m >= p['gov_start_month']:
            gov_revenue = p['gov_contract_value'] / 12   # 按月均摊

        # 3. 溯源平台分润（第12个月后且渠道打通）
        premium_revenue = 0.0
        if m >= 11:
            premium_landed = rng.random() < p['premium_prob']
            if premium_landed:
                premium_revenue = (active_clients
                                   * p['avg_acreage']
                                   * p['premium_per_mu_year']
                                   / 12)

        revenue = sub_revenue + gov_revenue + premium_revenue

        # ── 成本计算 ────────────────────────────────────────────────

        # 固定成本：折旧 + 云服务器 + 维护耗材
        fixed_cost = MONTHLY_FIXED_TOTAL

        # 变动成本：随客户数增长的售后人力
        variable_cost = active_clients * COST_PER_CLIENT

        cost = fixed_cost + variable_cost
        profit = revenue - cost
        cum_profit += profit

        monthly_revenue.append(revenue)
        monthly_cost.append(cost)
        monthly_clients.append(active_clients)
        cumulative_profit.append(cum_profit)
        monthly_profit.append(profit)

    # 回本月份（累计利润首次转正）
    breakeven_month = next(
        (i for i, v in enumerate(cumulative_profit) if v >= 0),
        None
    )

    return {
        'revenue':    np.array(monthly_revenue),
        'cost':       np.array(monthly_cost),
        'profit':     np.array(monthly_profit),
        'clients':    np.array(monthly_clients),
        'cum_profit': np.array(cumulative_profit),
        'breakeven':  breakeven_month,
        'final_profit': cumulative_profit[-1],
        'peak_clients': max(monthly_clients),
    }

# 运行所有情景
results = {sc: run_revenue_sim(sc) for sc in SCENARIOS}

# ══════════════════════════════════════════════════════════════════
# 绘图
# ══════════════════════════════════════════════════════════════════
fig = plt.figure(figsize=(16, 18))
fig.patch.set_facecolor('white')
gs = GridSpec(3, 3, figure=fig,
              hspace=0.42, wspace=0.32,
              top=0.93, bottom=0.05, left=0.08, right=0.97)

t = np.arange(MONTHS)

# 主标题
fig.text(0.5, 0.965,
         '果园守护者 · 四会落地盈收仿真（3年期）',
         ha='center', fontsize=15, fontweight='bold', color=SLATE)
fig.text(0.5, 0.940,
         f'初始投入 CAPEX={CAPEX/1e4:.1f}万元 | 轻资产运营 | 三种落地情景对比',
         ha='center', fontsize=10, color='#6b7280')

sc_keys   = list(SCENARIOS.keys())
sc_colors = [SCENARIOS[k]['color'] for k in sc_keys]
sc_ls     = [SCENARIOS[k]['ls']    for k in sc_keys]

# ── 第一行：月收入 + 月成本+利润 + 客户数 ───────────────────────

# 月收入曲线
ax_rev = fig.add_subplot(gs[0, :2])
for sc, color, ls in zip(sc_keys, sc_colors, sc_ls):
    res = results[sc]
    label = sc.replace('\n', ' ')
    ax_rev.plot(t, res['revenue'] / 1e4, color=color, ls=ls, lw=2, label=f'收入·{label}')

# 用基准情景的成本线作为参照
base = results['基准情景\n（正常推进）']
ax_rev.plot(t, base['cost'] / 1e4, color=MGRAY, lw=1.5, ls=':', label='月度成本（基准情景）')

# 基准情景盈亏区域填充
ax_rev.fill_between(t,
    base['cost'] / 1e4, base['revenue'] / 1e4,
    where=base['revenue'] >= base['cost'],
    alpha=0.09, color=FOREST, label='_nolegend_')
ax_rev.fill_between(t,
    base['cost'] / 1e4, base['revenue'] / 1e4,
    where=base['revenue'] < base['cost'],
    alpha=0.09, color=RED, label='_nolegend_')

# 标注政府合作窗口
ax_rev.axvline(x=9, color=MGRAY, lw=1, ls='--', alpha=0.5)
ax_rev.text(9.3, ax_rev.get_ylim()[1] * 0.05 if ax_rev.get_ylim()[1] > 0 else 0.5,
            '植保站\n合作窗口', fontsize=7.5, color=MGRAY)

ax_rev.set_xlabel('月份')
ax_rev.set_ylabel('万元 / 月')
ax_rev.set_title('月度收入 vs 成本（基准情景成本线为参照）', fontweight='bold')
ax_rev.legend(loc='upper left', fontsize=8, framealpha=0.9)
ax_rev.set_xlim(0, MONTHS - 1)
ax_rev.set_ylim(bottom=0)
ax_rev.grid(axis='y', alpha=0.3)

# 客户数增长
ax_cli = fig.add_subplot(gs[0, 2])
for sc, color, ls in zip(sc_keys, sc_colors, sc_ls):
    res = results[sc]
    ax_cli.plot(t, res['clients'], color=color, ls=ls, lw=2,
                label=sc.replace('\n', ' '))
ax_cli.set_xlabel('月份')
ax_cli.set_ylabel('付费合作社数（户）')
ax_cli.set_title('签约客户增长曲线', fontweight='bold')
ax_cli.legend(fontsize=8, framealpha=0.9)
ax_cli.set_xlim(0, MONTHS - 1)
ax_cli.set_ylim(bottom=0)
ax_cli.grid(axis='y', alpha=0.3)

# ── 第二行：月度净利润 + 累计利润曲线 + CAPEX 构成饼图 ──────────

# 月度净利润（柱状图）
ax_bar = fig.add_subplot(gs[1, :2])
width = 0.28
offsets = [-width, 0, width]
for (sc, color, ls), offset in zip(zip(sc_keys, sc_colors, sc_ls), offsets):
    res = results[sc]
    bars = ax_bar.bar(t + offset, res['profit'] / 1e4, width=width,
                      color=color, alpha=0.75, label=sc.replace('\n', ' '))

ax_bar.axhline(0, color=SLATE, lw=1.2, alpha=0.6)
ax_bar.set_xlabel('月份')
ax_bar.set_ylabel('万元 / 月')
ax_bar.set_title('月度净利润（正=盈利，负=亏损）', fontweight='bold')
ax_bar.legend(fontsize=8, framealpha=0.9)
ax_bar.set_xlim(-1, MONTHS)
ax_bar.grid(axis='y', alpha=0.3)

# CAPEX 构成饼图
ax_pie = fig.add_subplot(gs[1, 2])
pie_labels = [k.replace('×', 'x') for k in CAPEX_ITEMS.keys()]
pie_values = list(CAPEX_ITEMS.values())
pie_colors = ['#4a7c59', '#2d5a27', '#1a3a1a', '#7fba96']
wedges, texts, autotexts = ax_pie.pie(
    pie_values,
    labels=None,
    colors=pie_colors,
    autopct='%1.0f%%',
    startangle=140,
    pctdistance=0.75,
)
for at in autotexts:
    at.set_fontsize(10)
    at.set_color('white')
    at.set_fontweight('bold')

# 图例（放在饼图下方）
ax_pie.legend(wedges, [f'{l}\n{v/1e4:.1f}万' for l, v in zip(pie_labels, pie_values)],
              loc='lower center', bbox_to_anchor=(0.5, -0.35),
              fontsize=8, framealpha=0.9, ncol=1)
ax_pie.set_title(f'初始投入构成\n合计 {CAPEX/1e4:.1f} 万元', fontweight='bold')

# ── 第三行：累计利润曲线 + 敏感性分析 + 汇总指标表 ─────────────

# 累计利润
ax_cum = fig.add_subplot(gs[2, :2])
for sc, color, ls in zip(sc_keys, sc_colors, sc_ls):
    res = results[sc]
    label = sc.replace('\n', ' ')
    ax_cum.plot(t, res['cum_profit'] / 1e4, color=color, ls=ls, lw=2.2, label=label)

    # 标注回本时间
    be = res['breakeven']
    if be is not None and be < MONTHS:
        ax_cum.axvline(x=be, color=color, alpha=0.3, lw=1.2)
        ax_cum.annotate(
            f'M{be}\n回本',
            xy=(be, 0),
            xytext=(be + 0.8, res['cum_profit'].min() / 1e4 * 0.15),
            fontsize=8, color=color,
            arrowprops=dict(arrowstyle='->', color=color, lw=1),
        )

ax_cum.axhline(0, color=SLATE, lw=1.3, ls='-', alpha=0.5)
ax_cum.fill_between(t, 0,
    [min(v, 0) for v in base['cum_profit'] / 1e4],
    alpha=0.07, color=RED)
ax_cum.fill_between(t, 0,
    [max(v, 0) for v in base['cum_profit'] / 1e4],
    alpha=0.07, color=FOREST)
ax_cum.set_xlabel('月份')
ax_cum.set_ylabel('累计净利润（万元）')
ax_cum.set_title('累计利润曲线（初始投入 2.8 万已计入）', fontweight='bold')
ax_cum.legend(fontsize=9, framealpha=0.9)
ax_cum.set_xlim(0, MONTHS - 1)
ax_cum.grid(axis='y', alpha=0.3)

# 汇总指标表（用 matplotlib table）
ax_tbl = fig.add_subplot(gs[2, 2])
ax_tbl.axis('off')

def fmt_profit(v):
    if v >= 0:
        return f'+{v/1e4:.1f}万'
    return f'{v/1e4:.1f}万'

rows = []
for sc in sc_keys:
    res = results[sc]
    be  = res['breakeven']
    rows.append([
        sc.replace('\n', ''),
        f"M{be}" if be is not None else '未回本',
        fmt_profit(res['final_profit']),
        f"{res['peak_clients']:.0f}户",
    ])

col_labels = ['情景', '回本月份', '3年净利润', '峰值客户']
tab = ax_tbl.table(
    cellText=rows,
    colLabels=col_labels,
    cellLoc='center',
    loc='center',
    bbox=[0, 0.15, 1, 0.75],
)
tab.auto_set_font_size(False)
tab.set_fontsize(9.5)

# 表头颜色
header_colors = [SLATE, SLATE, SLATE, SLATE]
for col in range(4):
    tab[(0, col)].set_facecolor(SLATE)
    tab[(0, col)].set_text_props(color='white', fontweight='bold')
    tab[(0, col)].set_edgecolor('#d1d5db')

# 数据行颜色
row_bg = [
    ('#fee2e2', RED),     # 保守情景
    ('#e8f0e9', FOREST),  # 基准情景
    ('#e0f2f1', TEAL),    # 乐观情景
]
for row_idx, (bg, tc) in enumerate(row_bg):
    for col in range(4):
        tab[(row_idx + 1, col)].set_facecolor(bg)
        tab[(row_idx + 1, col)].set_edgecolor('#d1d5db')
        # 利润列：正绿负红
        if col == 2:
            val = rows[row_idx][2]
            if val.startswith('+'):
                tab[(row_idx + 1, col)].set_text_props(color='#166534', fontweight='bold')
            else:
                tab[(row_idx + 1, col)].set_text_props(color='#991b1b', fontweight='bold')

ax_tbl.set_title('盈亏关键指标汇总', fontweight='bold', fontsize=11, pad=10)

# 底部注释
ax_tbl.text(0.5, 0.08,
    '★ 保守情景：转化率25%，无政府合同\n'
    '★ 基准情景：转化率40%，40%概率签植保站\n'
    '★ 乐观情景：转化率55%，溯源渠道打通',
    transform=ax_tbl.transAxes,
    ha='center', va='bottom', fontsize=8,
    color='#374151', linespacing=1.6)

# ── 保存 ─────────────────────────────────────────────────────────
out_path = os.path.join(os.path.dirname(os.path.abspath(__file__)), 'sim_revenue.png')
fig.savefig(out_path, dpi=150, bbox_inches='tight', facecolor='white')
plt.close()
print(f'[完成] 盈收仿真图已保存至：{out_path}')

# ── 控制台摘要 ───────────────────────────────────────────────────
print('\n══ 盈收仿真结果摘要 ══')
print(f'  初始投入（CAPEX） ：{CAPEX/1e4:.1f} 万元')
print()
for sc in sc_keys:
    res = results[sc]
    be  = res['breakeven']
    label = sc.replace('\n', '')
    print(f'  {label}')
    print(f'    回本月份     ：{"第"+str(be)+"月" if be else "3年内未回本"}')
    print(f'    3年净利润    ：{fmt_profit(res["final_profit"])}')
    print(f'    峰值客户数   ：{res["peak_clients"]:.0f} 户')
    print()

print('══ 关键风险提示 ══')
print('  · 最大不确定因素：付费转化率（25%~55%），大于定价本身的影响')
print('  · 保守情景3年亏损，但亏损额有限（不超过初始投入的3倍），底线可控')
print('  · 政府/植保站采购是基准和乐观情景的关键增量，目标在第6-9月推进签约')
print('  · 溯源溢价分润在大创阶段不应作为必要收入，视为上行期权')
print('  · 建议：把第一年的KPI集中在"让示范点的果农愿意续费"而非快速扩张')