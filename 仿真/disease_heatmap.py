# -*- coding: utf-8 -*-
"""
果园守护者 · 四会专项
模块1：黄龙病 SI 网格扩散仿真
重点：每个网格独立感染概率热力图，含格线与数值标注
"""

import numpy as np
import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
import matplotlib.font_manager as fm
import matplotlib.ticker as ticker
from matplotlib.gridspec import GridSpec
from matplotlib.colors import LinearSegmentedColormap
from scipy.ndimage import gaussian_filter
import os, warnings
warnings.filterwarnings('ignore')

# ── 字体 ────────────────────────────────────────────────────────
def get_chn_font():
    cands = ['Microsoft YaHei','SimHei','PingFang SC',
             'Noto Sans CJK SC','DejaVu Sans']
    avail = {f.name for f in fm.fontManager.ttflist}
    for f in cands:
        if f in avail:
            return f
    return None

fn = get_chn_font()
if fn:
    plt.rcParams['font.family'] = fn
plt.rcParams.update({
    'axes.unicode_minus': False,
    'figure.dpi': 150,
    'axes.titlesize': 12,
    'axes.labelsize': 10,
    'xtick.labelsize': 8,
    'ytick.labelsize': 8,
})

# ── 自定义色图：健康绿→警戒黄→感染红 ──────────────────────────
CMAP = LinearSegmentedColormap.from_list('hlb', [
    '#1a5c2a',   # 0.0  深绿：健康
    '#4a9e5c',   # 0.2  中绿
    '#a8d5a2',   # 0.4  浅绿
    '#f5e642',   # 0.55 黄：警戒
    '#e87b2a',   # 0.70 橙：高风险
    '#c0392b',   # 0.85 红：感染
    '#7b0d0d',   # 1.0  深红：重度感染
])

# ══════════════════════════════════════════════════════════════════
# 仿真参数与函数（与原版一致）
# ══════════════════════════════════════════════════════════════════
GRID = 20     # 改为20×20，方便在图上显示每格数值
DAYS = 365
np.random.seed(42)

def make_terrain(n):
    return gaussian_filter(np.random.randn(n, n), sigma=4) * 2.5

def init_infection(n, seeds=4):
    I = np.zeros((n, n))
    for _ in range(seeds):
        r, c = np.random.randint(2, n-2, 2)
        I[r, c] = 1.0
    return I

def wind_bias_weights(theta_wind=np.radians(135), alpha=0.8):
    dirs = [(-1,-1),(-1,0),(-1,1),(0,-1),(0,1),(1,-1),(1,0),(1,1)]
    W = {}
    for dr, dc in dirs:
        theta_ij = np.arctan2(dr, dc)
        bias = max(1 + alpha * np.cos(theta_ij - theta_wind), 0.05)
        W[(dr, dc)] = bias
    return W

def beta_t(day):
    month = (day % 365) // 30 + 1
    return 0.10 if month in [4,5,6,9,10] else 0.02

def run_SI(I0, gamma, apply_isolation=False, n=GRID, days=DAYS,
           snapshot_days=None):
    """
    仿真SI扩散，返回逐天全园感染率 + 指定天数的网格快照

    snapshot_days: list，需要保存网格状态的天数列表
    """
    S = 1.0 - I0.copy()
    I = I0.copy()
    W = wind_bias_weights()
    sigma_g = 1.5

    history   = []
    snapshots = {}   # {day: I矩阵}

    for day in range(days):
        beta = beta_t(day)
        if apply_isolation:
            beta *= 0.12

        pressure = np.zeros((n, n))
        for (dr, dc), w_bias in W.items():
            shifted = np.roll(np.roll(I, dr, axis=0), dc, axis=1)
            dist = np.sqrt(dr**2 + dc**2)
            w = np.exp(-dist**2 / (2 * sigma_g**2)) * w_bias
            pressure += beta * S * shifted * w

        new_I = np.clip(I + pressure - gamma * I, 0, 1)
        new_S = np.clip(S - pressure + gamma * I, 0, 1)
        I, S  = new_I, new_S
        history.append(I.mean())

        if snapshot_days and day in snapshot_days:
            snapshots[day] = I.copy()

    return np.array(history), I, snapshots

# ══════════════════════════════════════════════════════════════════
# 核心绘图函数：单张网格热力图
# ══════════════════════════════════════════════════════════════════
def draw_heatmap(ax, grid_data, title, title_color,
                 show_values=True, grid_alpha=0.35, fontsize_val=5.5):
    """
    画一张真正的"格子热力图"
    - 每格颜色 = 该格感染率
    - 每格中心标注感染率数值（高感染率格子用白字，低感染率用深字）
    - 格线清晰可见
    - 坐标轴显示网格编号（对应果园坐标系）
    """
    n = grid_data.shape[0]
    mean_val = grid_data.mean()

    im = ax.imshow(grid_data, cmap=CMAP, vmin=0, vmax=1,
                   interpolation='nearest', aspect='equal')

    # ── 网格线 ──────────────────────────────────────────────────
    # 用minor tick画网格线，比直接用gridlines更精确
    ax.set_xticks(np.arange(-0.5, n, 1), minor=True)
    ax.set_yticks(np.arange(-0.5, n, 1), minor=True)
    ax.grid(which='minor', color='white', linewidth=0.6, alpha=grid_alpha)
    ax.tick_params(which='minor', bottom=False, left=False)

    # 主刻度：每5格一个编号（坐标轴标注网格ID）
    major_ticks = np.arange(0, n, 5)
    ax.set_xticks(major_ticks)
    ax.set_yticks(major_ticks)
    ax.set_xticklabels([f'{i*10}m' for i in major_ticks], fontsize=7)
    ax.set_yticklabels([f'{i*10}m' for i in major_ticks], fontsize=7)

    # ── 每格数值标注 ─────────────────────────────────────────────
    if show_values:
        for r in range(n):
            for c in range(n):
                val = grid_data[r, c]
                if val < 0.03:
                    continue   # 健康格不标（太密），只标有感染迹象的格子
                # 颜色：感染率>0.5用白字，否则用深绿字（确保可读性）
                txt_color = 'white' if val > 0.45 else '#1a3a1a'
                txt = f'{val:.0%}' if val >= 0.10 else f'{val:.1%}'
                ax.text(c, r, txt,
                        ha='center', va='center',
                        fontsize=fontsize_val, color=txt_color,
                        fontweight='bold' if val > 0.5 else 'normal')

    # ── 标题与感染率统计 ─────────────────────────────────────────
    # 统计感染格分布
    infected_hi  = (grid_data > 0.7).sum()   # 重度感染（>70%）
    infected_mid = ((grid_data > 0.3) & (grid_data <= 0.7)).sum()  # 中度
    infected_lo  = ((grid_data > 0.05) & (grid_data <= 0.3)).sum() # 轻度
    healthy      = (grid_data <= 0.05).sum()                        # 健康

    ax.set_title(title, fontweight='bold', color=title_color,
                 fontsize=11, pad=6)

    # 底部信息条（感染分布统计）
    info = (f'均值 {mean_val:.1%}  |  '
            f'重度 {infected_hi}格  '
            f'中度 {infected_mid}格  '
            f'轻度 {infected_lo}格  '
            f'健康 {healthy}格')
    ax.set_xlabel(info, fontsize=7.5, color='#374151', labelpad=4)

    return im

# ══════════════════════════════════════════════════════════════════
# 运行仿真 —— 额外保存 Day90/180/270 中间快照
# ══════════════════════════════════════════════════════════════════
I0      = init_infection(GRID, seeds=4)
terrain = make_terrain(GRID)

SNAP_DAYS = [89, 179, 269, 364]   # Day90/180/270/365（index从0）

hist_trad, map_trad, snap_trad = run_SI(
    I0, gamma=0.02, apply_isolation=False, snapshot_days=SNAP_DAYS)
hist_lo,   map_lo,   snap_lo   = run_SI(
    I0, gamma=0.08, apply_isolation=True,  snapshot_days=SNAP_DAYS)
hist_hi,   map_hi,   snap_hi   = run_SI(
    I0, gamma=0.15, apply_isolation=True,  snapshot_days=SNAP_DAYS)

# ══════════════════════════════════════════════════════════════════
# 图1：扩散曲线（独立保存，宽图）
# ══════════════════════════════════════════════════════════════════
fig_curve, ax_c = plt.subplots(figsize=(14, 5))
fig_curve.patch.set_facecolor('white')

t = np.arange(DAYS)
ax_c.plot(t, hist_trad*100, color='#c0392b', lw=2.2, label='传统人工防控  γ=0.02')
ax_c.plot(t, hist_lo*100,   color='#d4870a', lw=2.0, ls='--',
          label='系统磨合期  γ=0.08（隔离带启用）')
ax_c.plot(t, hist_hi*100,   color='#2d5a27', lw=2.2,
          label='系统稳定期  γ=0.15（隔离带启用）')

for ms, me, lbl in [(3,5,'春梢期'), (8,9,'秋梢期')]:
    ax_c.axvspan(ms*30, me*30, alpha=0.08, color='#d4870a')
    ax_c.text((ms+me)/2*30, 92, lbl, ha='center', fontsize=9, color='#b45309')

# 快照时间线
for d, lbl in zip(SNAP_DAYS, ['Day90','Day180','Day270','Day365']):
    ax_c.axvline(x=d, color='#9ca3af', lw=0.8, ls=':', alpha=0.7)
    ax_c.text(d+2, 2, lbl, fontsize=7.5, color='#6b7280')

ax_c.set_xlabel('天（第1年）')
ax_c.set_ylabel('全园平均感染率（%）')
ax_c.set_title('黄龙病 SI 扩散曲线 — 三情景对比（20×20网格，≈40亩）',
               fontweight='bold', fontsize=13)
ax_c.legend(loc='upper left', framealpha=0.9)
ax_c.set_xlim(0, DAYS); ax_c.set_ylim(0, 100)
ax_c.grid(axis='y', alpha=0.3)
ax_c.spines['top'].set_visible(False)
ax_c.spines['right'].set_visible(False)

out1 = os.path.join(os.path.dirname(os.path.abspath(__file__)), 'sim_curve.png')
fig_curve.savefig(out1, dpi=150, bbox_inches='tight', facecolor='white')
plt.close(fig_curve)
print(f'[完成] 扩散曲线图 → {out1}')

# ══════════════════════════════════════════════════════════════════
# 图2：4列×3行 网格热力图（Day90/180/270/365 × 三情景）
# ══════════════════════════════════════════════════════════════════
COLS = 4   # 时间点
ROWS = 3   # 情景

fig_maps = plt.figure(figsize=(20, 14))
fig_maps.patch.set_facecolor('white')
gs = GridSpec(ROWS, COLS, figure=fig_maps,
              hspace=0.55, wspace=0.18,
              top=0.91, bottom=0.05, left=0.04, right=0.95)

fig_maps.text(0.5, 0.955,
    '黄龙病扩散热力图 — 各网格感染概率（20×20格，每格10m×10m）',
    ha='center', fontsize=14, fontweight='bold', color='#1a2e1a')
fig_maps.text(0.5, 0.930,
    '数值 = 该网格感染率 | 仅标注感染率>3%的格子 | 格线为单个10m×10m防控单元',
    ha='center', fontsize=9.5, color='#6b7280')

scenario_info = [
    ('传统人工防控', '#c0392b', snap_trad),
    ('系统磨合期',   '#d4870a', snap_lo),
    ('系统稳定期',   '#2d5a27', snap_hi),
]
day_labels = ['Day 90（3个月）', 'Day 180（半年）',
              'Day 270（9个月）','Day 365（满1年）']

# 行标签（情景）
for row, (sc_name, sc_color, _) in enumerate(scenario_info):
    fig_maps.text(0.005, 1 - (row + 0.5) / ROWS * 0.88 - 0.06,
                  sc_name, va='center', ha='left',
                  fontsize=11, fontweight='bold', color=sc_color,
                  rotation=90)

# 列标签（时间）已在每个子图标题里

ims = []
for row, (sc_name, sc_color, snaps) in enumerate(scenario_info):
    for col, (day_idx, day_lbl) in enumerate(zip(SNAP_DAYS, day_labels)):
        ax = fig_maps.add_subplot(gs[row, col])
        data = snaps[day_idx]

        # 第一行加列标题
        title = day_lbl if row == 0 else ''
        # 每格都显示情景+时间组合标题（紧凑版）
        cell_title = f'{day_lbl}' if row == 0 else f'({data.mean():.1%})'

        im = draw_heatmap(ax, data,
                          title=cell_title,
                          title_color=sc_color if row > 0 else '#374151',
                          show_values=True,
                          fontsize_val=5.2)
        ims.append(im)

        # 左侧情景标注（每行第一列）
        if col == 0:
            ax.set_ylabel(sc_name, fontsize=10, fontweight='bold',
                          color=sc_color, labelpad=6)

# 色条（共享）
cbar_ax = fig_maps.add_axes([0.956, 0.08, 0.012, 0.80])
sm = plt.cm.ScalarMappable(cmap=CMAP,
                            norm=plt.Normalize(vmin=0, vmax=1))
sm.set_array([])
cbar = fig_maps.colorbar(sm, cax=cbar_ax)
cbar.set_label('单格感染概率', fontsize=10, labelpad=8)
cbar.set_ticks([0, 0.2, 0.4, 0.6, 0.8, 1.0])
cbar.set_ticklabels(['0%\n健康', '20%', '40%\n警戒',
                     '60%', '80%\n高危', '100%\n感染'])
cbar.ax.tick_params(labelsize=8)

out2 = os.path.join(os.path.dirname(os.path.abspath(__file__)), 'sim_heatmaps.png')
fig_maps.savefig(out2, dpi=150, bbox_inches='tight', facecolor='white')
plt.close(fig_maps)
print(f'[完成] 网格热力图    → {out2}')

# ══════════════════════════════════════════════════════════════════
# 图3：初始状态 vs 四种时间切片对比（针对传统防控，最直观）
# ══════════════════════════════════════════════════════════════════
fig_trad = plt.figure(figsize=(20, 5))
fig_trad.patch.set_facecolor('white')
gs3 = GridSpec(1, 5, figure=fig_trad,
               hspace=0.3, wspace=0.18,
               top=0.85, bottom=0.12, left=0.04, right=0.94)

fig_trad.text(0.5, 0.95,
    '传统防控情景：感染扩散时序（每格感染率，粤西北东南季风方向）',
    ha='center', fontsize=13, fontweight='bold', color='#c0392b')

all_snaps = [('初始状态\nDay 0', I0)]
for d, lbl in zip(SNAP_DAYS, ['Day 90', 'Day 180', 'Day 270', 'Day 365']):
    all_snaps.append((lbl, snap_trad[d]))

for col, (lbl, data) in enumerate(all_snaps):
    ax = fig_trad.add_subplot(gs3[0, col])
    draw_heatmap(ax, data, title=lbl, title_color='#c0392b',
                 show_values=True, fontsize_val=5.0)
    # 箭头连接（除最后一列）
    if col < 4:
        fig_trad.text(
            0.04 + (col + 0.92) / 5 * 0.90,
            0.50, '→', ha='center', va='center',
            fontsize=18, color='#9ca3af',
            transform=fig_trad.transFigure
        )

cbar_ax3 = fig_trad.add_axes([0.945, 0.12, 0.01, 0.70])
sm3 = plt.cm.ScalarMappable(cmap=CMAP, norm=plt.Normalize(0,1))
sm3.set_array([])
cb3 = fig_trad.colorbar(sm3, cax=cbar_ax3)
cb3.set_ticks([0, 0.5, 1.0])
cb3.set_ticklabels(['0%\n健康', '50%', '100%\n感染'])
cb3.ax.tick_params(labelsize=8)

out3 = os.path.join(os.path.dirname(os.path.abspath(__file__)), 'sim_spread_seq.png')
fig_trad.savefig(out3, dpi=150, bbox_inches='tight', facecolor='white')
plt.close(fig_trad)
print(f'[完成] 传统防控时序图 → {out3}')

# ── 控制台摘要 ──────────────────────────────────────────────────
print('\n══ 结果摘要 ══')
print(f'  初始感染率        ：{I0.mean()*100:.1f}%')
print(f'  传统防控 Day365   ：{map_trad.mean()*100:.1f}%')
print(f'  系统磨合 Day365   ：{map_lo.mean()*100:.2f}%')
print(f'  系统稳定 Day365   ：{map_hi.mean()*100:.2f}%')
print(f'\n  输出文件：')
print(f'    {out1}')
print(f'    {out2}')
print(f'    {out3}')