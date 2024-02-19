(() => {
    'use strict';
    /*!
     * Color mode toggler for Bootstrap's docs (https://getbootstrap.com/)
     * Copyright 2011-2023 The Bootstrap Authors
     * Licensed under the Creative Commons Attribution 3.0 Unported License.
     */
    function toggleTheme() {
        // 设置主题
        const setTheme = function (theme) {
            if (theme === 'auto' && window.matchMedia('(prefers-color-scheme: dark)').matches) {
                document.documentElement.setAttribute('data-bs-theme', 'dark')
            } else {
                document.documentElement.setAttribute('data-bs-theme', theme)
            }
        }
        // 获取主题偏好
        const getPreferredTheme = () => {
            const storedTheme = localStorage.getItem('theme')
            if (storedTheme) {
                return storedTheme
            }
            return 'dark'
        }
        // 立即设置一次主题
        setTheme(getPreferredTheme())

        // 监听系统开启夜间模式
        window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', () => {
            if (localStorage.getItem('theme') === 'auto') {
                setTheme(getPreferredTheme())
            }
        })

        // 切换主题下拉按钮旁边显示当前主题
        const showActiveTheme = (theme, focus = false) => {
            const themeSwitcher = document.querySelector('#bd-theme')
            if (!themeSwitcher) {
                return
            }

            const themeSwitcherText = document.querySelector('#bd-theme-text')
            const activeThemeIcon = document.querySelector('.theme-icon-active')
            const btnToActive = document.querySelector(`[data-bs-theme-value="${theme}"]`)
            const activeIcon = btnToActive.querySelector('i.theme-icon')

            document.querySelectorAll('[data-bs-theme-value]').forEach(element => {
                element.classList.remove('active')
                element.setAttribute('aria-pressed', 'false')
            })
            btnToActive.classList.add('active')
            btnToActive.setAttribute('aria-pressed', 'true')

            const classList = ([...activeThemeIcon.classList]).filter(s => !s.startsWith('bi-'))
            const activeClass = [...activeIcon.classList].find(s => s.startsWith('bi-'))
            classList.push(activeClass)
            activeThemeIcon.className = classList.join(' ')

            const themeSwitcherLabel = `${themeSwitcherText.textContent} (${btnToActive.innerText.trim()})`
            themeSwitcher.setAttribute('aria-label', themeSwitcherLabel)

            if (focus) {
                themeSwitcher.focus()
            }
        }

        // 页面加载后执行
        window.addEventListener('DOMContentLoaded', () => {
            showActiveTheme(getPreferredTheme())

            // 手动点击主题切换按钮
            document.querySelectorAll('[data-bs-theme-value]')
                .forEach(btn => {
                    btn.addEventListener('click', () => {
                        const theme = btn.getAttribute('data-bs-theme-value')
                        localStorage.setItem('theme', theme)
                        setTheme(theme)
                        showActiveTheme(theme, true)
                    })
                })
        })
    }
    toggleTheme()

    function login() {
        const form = document.getElementById('login-form')
        if (!form) { return }
        const fieldset = document.getElementById('fieldset')
        fieldset.disabled = false
        const plain = document.getElementById('plain')
        const passwd = document.getElementById('passwd')
        form.addEventListener('submit', e => {
            passwd.value = sha256(form.username.value + sha256(plain.value) + form.salt.value)
            return true
        })
    }
    login()

})()