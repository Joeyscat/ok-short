from selenium import webdriver
from time import sleep

from selenium.webdriver.common.keys import Keys


class Oook(object):
    def __init__(self):
        self.driver = webdriver.Chrome()
        self.driver.get('http://link.oook.fun/')

    def position(self):
        test_link = 'https://www.w3school.com.cn/html/html_attributes.asp'
        ele = self.driver.find_element_by_class_name('el-input__inner')
        ele.send_keys(test_link)
        btn = self.driver.find_element_by_xpath('//*[@id="app"]/div[2]/div/div[1]/div/div[2]/div[2]/button/span')
        btn.click()
        sleep(1)
        link_span = self.driver.find_element_by_xpath('//*[@id="app"]/div[2]/div/div[2]/div/div/div[1]/div[1]/a/span')
        link_content = link_span.get_attribute('textContent')
        print(link_content)
        short_link = link_content.split('：')[1].split(' ')[0]
        print(short_link)
        sleep(1)
        webdriver.Chrome().get('http://' + short_link)
        sleep(3)
        self.driver.switch_to.window(self.driver.window_handles[0])  # 切换当前页面标签
        new_link = self.driver.current_url
        print(new_link)
        if new_link == test_link:
            print('恭喜你，测试成功！')
        else:
            print('很遗憾，测试失败，请继续加油！')
        self.driver.quit()


if __name__ == '__main__':
    case = Oook()
    case.position()
