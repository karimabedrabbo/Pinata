//
//  Extensions.swift
//  youtube
//
//  Created by Brian Voong on 6/3/16.
//  Copyright © 2016 letsbuildthatapp. All rights reserved.
//

import UIKit
import Foundation
import AVFoundation

extension AVAsset {

    func videoOrientation() -> UIInterfaceOrientation {
        
        let size = tracks(withMediaType: AVMediaType.video).first?.naturalSize ?? .zero
        
        guard let transform = tracks(withMediaType: AVMediaType.video).first?.preferredTransform else {
            return .portrait
        }
        
        switch (transform.tx, transform.ty) {
        case (0, 0):
            return .landscapeRight
        case (size.width, size.height):
            return .landscapeLeft
        case (0, size.width):
            return .portraitUpsideDown
        default:
            return .portrait
        }
    }
}



// generic method, can be in a category
extension CGPoint {
    func distance(to point: CGPoint) -> CGFloat {
        // there is already a function for sqrt(x * x + y * y)
        return hypot(self.x - point.x, self.y - point.y)
    }
}



extension UIView {
    /**
     Convert UIView to UIImage
     */
    func toImage() -> UIImage {
        UIGraphicsBeginImageContextWithOptions(self.bounds.size, !self.isOpaque, 0.0)
        self.drawHierarchy(in: self.bounds, afterScreenUpdates: false)
        let snapshotImageFromMyView = UIGraphicsGetImageFromCurrentImageContext()
        UIGraphicsEndImageContext()
        return snapshotImageFromMyView!
    }
}

extension UIImageView {
    
    func alphaAtPoint(_ point: CGPoint) -> CGFloat {
        
        var pixel: [UInt8] = [0, 0, 0, 0]
        let colorSpace = CGColorSpaceCreateDeviceRGB();
        let alphaInfo = CGImageAlphaInfo.premultipliedLast.rawValue
        
        guard let context = CGContext(data: &pixel, width: 1, height: 1, bitsPerComponent: 8, bytesPerRow: 4, space: colorSpace, bitmapInfo: alphaInfo) else {
            return 0
        }
        
        context.translateBy(x: -point.x, y: -point.y);
        
        layer.render(in: context)
        
        let floatAlpha = CGFloat(pixel[3])
        
        return floatAlpha
    }
    
}




extension UIColor {
    /**
     Create a ligher color
     */
    func lighterHsb(by percentage: CGFloat = 30.0) -> UIColor {
        return self.adjustBrightnessHsb(by: abs(percentage))
    }
    
    /**
     Create a darker color
     */
    func darkerHsb(by percentage: CGFloat = 30.0) -> UIColor {
        return self.adjustBrightnessHsb(by: -abs(percentage))
    }
    

    /**
     Try to increase brightness or decrease saturation
     */
    func adjustBrightnessHsb(by percentage: CGFloat = 30.0) -> UIColor {
        var h: CGFloat = 0, s: CGFloat = 0, b: CGFloat = 0, a: CGFloat = 0
        if self.getHue(&h, saturation: &s, brightness: &b, alpha: &a) {
                let newB: CGFloat = min(max(b + (percentage/100.0)*b, 0.0), 1.0)
                return UIColor(hue: h, saturation: s, brightness: newB, alpha: a)
        }
        return self
    }
}

extension UIImage {
    
    func overlayWith(image: UIImage, posX: CGFloat, posY: CGFloat) -> UIImage {
        let newWidth = size.width < posX + image.size.width ? posX + image.size.width : size.width
        let newHeight = size.height < posY + image.size.height ? posY + image.size.height : size.height
        let newSize = CGSize(width: newWidth, height: newHeight)
        
        UIGraphicsBeginImageContextWithOptions(newSize, false, 0.0)
        draw(in: CGRect(origin: CGPoint.zero, size: size))
        image.draw(in: CGRect(origin: CGPoint(x: posX, y: posY), size: image.size))
        let newImage = UIGraphicsGetImageFromCurrentImageContext()!
        UIGraphicsEndImageContext()
        
        return newImage
    }
    
}





